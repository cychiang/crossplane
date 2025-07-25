/*
Copyright 2023 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package usage

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/crossplane-runtime/pkg/test"

	"github.com/crossplane/crossplane/apis/protection/v1beta1"
	"github.com/crossplane/crossplane/internal/protection"
	"github.com/crossplane/crossplane/internal/protection/usage"
)

var _ admission.Handler = &Handler{}

type FinderFn func(ctx context.Context, o usage.Object) ([]protection.Usage, error)

func (fn FinderFn) FindUsageOf(ctx context.Context, o usage.Object) ([]protection.Usage, error) {
	return fn(ctx, o)
}

var errBoom = errors.New("boom")

func TestHandle(t *testing.T) {
	protected := "This resource is protected!"

	type params struct {
		client client.Client
		f      Finder
		opts   []HandlerOption
	}

	type args struct {
		request admission.Request
	}

	type want struct {
		resp admission.Response
	}

	cases := map[string]struct {
		reason string
		params params
		args   args
		want   want
	}{
		"UnexpectedCreate": {
			reason: "We should return an error if the request is a create (not a delete).",
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Create,
					},
				},
			},
			want: want{
				resp: admission.Errored(http.StatusBadRequest, errors.Errorf(errFmtUnexpectedOp, admissionv1.Create)),
			},
		},
		"UnexpectedConnect": {
			reason: "We should return an error if the request is a connect (not a delete).",
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Connect,
					},
				},
			},
			want: want{
				resp: admission.Errored(http.StatusBadRequest, errors.Errorf(errFmtUnexpectedOp, admissionv1.Connect)),
			},
		},
		"UnexpectedUpdate": {
			reason: "We should return an error if the request is an update (not a delete).",
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Update,
					},
				},
			},
			want: want{
				resp: admission.Errored(http.StatusBadRequest, errors.Errorf(errFmtUnexpectedOp, admissionv1.Update)),
			},
		},
		"UnexpectedOperation": {
			reason: "We should return an error if the request is unknown (not a delete).",
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Operation("unknown"),
					},
				},
			},
			want: want{
				resp: admission.Errored(http.StatusBadRequest, errors.Errorf(errFmtUnexpectedOp, admissionv1.Operation("unknown"))),
			},
		},
		"DeleteWithoutOldObj": {
			reason: "We should not return an error if delete request does not have the old object.",
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Delete,
					},
				},
			},
			want: want{
				resp: admission.Errored(http.StatusBadRequest, errors.New("unexpected end of JSON input")),
			},
		},
		"DeleteAllowedNoUsages": {
			reason: "We should allow a delete request if there is no usages for the given object.",
			params: params{
				f: FinderFn(func(_ context.Context, _ usage.Object) ([]protection.Usage, error) {
					return nil, nil
				}),
			},
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Delete,
						OldObject: runtime.RawExtension{
							Raw: []byte(`{
								"apiVersion": "nop.crossplane.io/v1alpha1",
								"kind": "NopResource",
								"metadata": {
									"name": "used-resource"
								}}`),
						},
					},
				},
			},
			want: want{
				resp: admission.Allowed(""),
			},
		},
		"DeleteRejectedCannotList": {
			reason: "We should reject a delete request if we cannot list usages.",
			params: params{
				f: FinderFn(func(_ context.Context, _ usage.Object) ([]protection.Usage, error) {
					return nil, errBoom
				}),
			},
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Delete,
						OldObject: runtime.RawExtension{
							Raw: []byte(`{
								"apiVersion": "nop.crossplane.io/v1alpha1",
								"kind": "NopResource",
								"metadata": {
									"name": "used-resource"
								}}`),
						},
					},
				},
			},
			want: want{
				resp: admission.Errored(http.StatusInternalServerError, errBoom),
			},
		},
		"DeleteBlockedWithUsageBy": {
			reason: "We should reject a delete request if there are usages for the given object with \"by\" defined.",
			params: params{
				client: &test.MockClient{
					MockPatch: func(_ context.Context, _ client.Object, _ client.Patch, _ ...client.PatchOption) error {
						return nil
					},
				},
				f: FinderFn(func(_ context.Context, _ usage.Object) ([]protection.Usage, error) {
					usages := []protection.Usage{
						&v1beta1.Usage{
							ObjectMeta: metav1.ObjectMeta{
								Namespace: "default",
								Name:      "used-by-some-resource",
							},
							Spec: v1beta1.UsageSpec{
								Of: v1beta1.NamespacedResource{
									APIVersion: "nop.crossplane.io/v1alpha1",
									Kind:       "NopResource",
									ResourceRef: &v1beta1.NamespacedResourceRef{
										Name: "used-resource",
									},
								},
								By: &v1beta1.Resource{
									APIVersion: "nop.crossplane.io/v1alpha1",
									Kind:       "NopResource",
									ResourceRef: &v1beta1.ResourceRef{
										Name: "using-resource",
									},
								},
							},
						},
					}
					return usages, nil
				}),
			},
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Delete,
						OldObject: runtime.RawExtension{
							Raw: []byte(`{
								"apiVersion": "nop.crossplane.io/v1alpha1",
								"kind": "NopResource",
								"metadata": {
									"name": "used-resource"
								}}`),
						},
					},
				},
			},
			want: want{
				resp: admission.Response{
					AdmissionResponse: admissionv1.AdmissionResponse{
						Allowed: false,
						Result: &metav1.Status{
							Code:   int32(http.StatusConflict),
							Reason: metav1.StatusReason("This resource is in-use by 1 usage(s), including the *v1beta1.Usage \"used-by-some-resource\" (in namespace \"default\") by resource NopResource/using-resource."),
						},
					},
				},
			},
		},
		"DeleteBlockedWithUsageReason": {
			reason: "We should reject a delete request if there are usages for the given object with \"reason\" defined.",
			params: params{
				client: &test.MockClient{
					MockPatch: func(_ context.Context, _ client.Object, _ client.Patch, _ ...client.PatchOption) error {
						return nil
					},
				},
				f: FinderFn(func(_ context.Context, _ usage.Object) ([]protection.Usage, error) {
					usages := []protection.Usage{
						&v1beta1.Usage{
							ObjectMeta: metav1.ObjectMeta{
								Name: "used-by-some-resource",
							},
							Spec: v1beta1.UsageSpec{
								Of: v1beta1.NamespacedResource{
									APIVersion: "nop.crossplane.io/v1alpha1",
									Kind:       "NopResource",
									ResourceRef: &v1beta1.NamespacedResourceRef{
										Name: "used-resource",
									},
								},
								Reason: &protected,
							},
						},
					}
					return usages, nil
				}),
			},
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Delete,
						OldObject: runtime.RawExtension{
							Raw: []byte(`{
								"apiVersion": "nop.crossplane.io/v1alpha1",
								"kind": "NopResource",
								"metadata": {
									"name": "used-resource"
								}}`),
						},
					},
				},
			},
			want: want{
				resp: admission.Response{
					AdmissionResponse: admissionv1.AdmissionResponse{
						Allowed: false,
						Result: &metav1.Status{
							Code:   int32(http.StatusConflict),
							Reason: metav1.StatusReason("This resource is in-use by 1 usage(s), including the *v1beta1.Usage \"used-by-some-resource\" with reason: \"This resource is protected!\"."),
						},
					},
				},
			},
		},
		"DeleteBlockedWithUsageNone": {
			reason: "We should reject a delete request if there are usages for the given object without \"reason\" or \"by\" defined.",
			params: params{
				client: &test.MockClient{
					MockPatch: func(_ context.Context, _ client.Object, _ client.Patch, _ ...client.PatchOption) error {
						return nil
					},
				},
				f: FinderFn(func(_ context.Context, _ usage.Object) ([]protection.Usage, error) {
					usages := []protection.Usage{
						&v1beta1.Usage{
							ObjectMeta: metav1.ObjectMeta{
								Name: "used-by-some-resource",
							},
							Spec: v1beta1.UsageSpec{
								Of: v1beta1.NamespacedResource{
									APIVersion: "nop.crossplane.io/v1alpha1",
									Kind:       "NopResource",
									ResourceRef: &v1beta1.NamespacedResourceRef{
										Name: "used-resource",
									},
								},
							},
						},
					}
					return usages, nil
				}),
			},
			args: args{
				request: admission.Request{
					AdmissionRequest: admissionv1.AdmissionRequest{
						Operation: admissionv1.Delete,
						OldObject: runtime.RawExtension{
							Raw: []byte(`{
								"apiVersion": "nop.crossplane.io/v1alpha1",
								"kind": "NopResource",
								"metadata": {
									"name": "used-resource"
								}}`),
						},
					},
				},
			},
			want: want{
				resp: admission.Response{
					AdmissionResponse: admissionv1.AdmissionResponse{
						Allowed: false,
						Result: &metav1.Status{
							Code:   int32(http.StatusConflict),
							Reason: metav1.StatusReason("This resource is in-use by 1 usage(s), including the *v1beta1.Usage \"used-by-some-resource\"."),
						},
					},
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			h := NewHandler(tc.params.client, tc.params.f, tc.params.opts...)

			got := h.Handle(context.Background(), tc.args.request)
			if diff := cmp.Diff(tc.want.resp, got); diff != "" {
				t.Errorf("%s\nHandle(...): -want response, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
