package revision

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestDeploymentWithRuntimeContainer(t *testing.T) {
	type args struct {
		deployment *appsv1.Deployment
	}
	type want struct {
		deployment *appsv1.Deployment
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"NoContainers": {
			reason: "Should not add the runtime container if there are no containers",
			args: args{
				deployment: &appsv1.Deployment{},
			},
			want: want{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: runtimeContainerName,
									},
								},
							},
						},
					},
				},
			},
		},
		"AlreadyFirstAndOnlyContainer": {
			reason: "Should do nothing if the runtime container is already the first and only container",
			args: args{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: runtimeContainerName,
									},
								},
							},
						},
					},
				},
			},
			want: want{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: runtimeContainerName,
									},
								},
							},
						},
					},
				},
			},
		},
		"AddedToExistingContainers": {
			reason: "Should not add the container to the existing containers as the first container",
			args: args{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: "some-other-container",
									},
								},
							},
						},
					},
				},
			},
			want: want{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: runtimeContainerName,
									},
									{
										Name: "some-other-container",
									},
								},
							},
						},
					},
				},
			},
		},
		"ExistButInWrongPlace": {
			reason: "Should move the runtime container to the first container position if it exists but is not the first container",
			args: args{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: "some-other-container",
									},
									{
										Name: runtimeContainerName,
									},
									{
										Name: "another-one",
									},
								},
							},
						},
					},
				},
			},
			want: want{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: runtimeContainerName,
									},
									{
										Name: "some-other-container",
									},
									{
										Name: "another-one",
									},
								},
							},
						},
					},
				},
			},
		},
		"ExistAtTheEnd": {
			reason: "Should move the runtime container to the first container position if it exists as the last container",
			args: args{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: "some-other-container",
									},
									{
										Name: "another-one",
									},
									{
										Name: runtimeContainerName,
									},
								},
							},
						},
					},
				},
			},
			want: want{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name: runtimeContainerName,
									},
									{
										Name: "some-other-container",
									},
									{
										Name: "another-one",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			DeploymentWithRuntimeContainer()(tc.args.deployment)
			if diff := cmp.Diff(tc.want.deployment, tc.args.deployment); diff != "" {
				t.Errorf("\n%s\nDeploymentWithRuntimeContainer(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestDeploymentRuntimeWithAdditionalPorts(t *testing.T) {
	type args struct {
		deployment *appsv1.Deployment
		ports      []corev1.ContainerPort
	}
	type want struct {
		deployment *appsv1.Deployment
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"NoPorts": {
			reason: "Should add the given ports if no ports are set",
			args: args{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{},
								},
							},
						},
					},
				},
				ports: []corev1.ContainerPort{
					{ContainerPort: 80, Name: "http"},
					{ContainerPort: 443, Name: "https"},
				},
			},
			want: want{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Ports: []corev1.ContainerPort{
											{ContainerPort: 80, Name: "http"},
											{ContainerPort: 443, Name: "https"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"DontOverridePorts": {
			reason: "Should add only new ports and not override existing ports",
			args: args{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Ports: []corev1.ContainerPort{
											{ContainerPort: 8080, Name: "http"},
										},
									},
								},
							},
						},
					},
				},
				ports: []corev1.ContainerPort{
					{ContainerPort: 80, Name: "http"},
					{ContainerPort: 443, Name: "https"},
				},
			},
			want: want{
				deployment: &appsv1.Deployment{
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Ports: []corev1.ContainerPort{
											{ContainerPort: 8080, Name: "http"},
											{ContainerPort: 443, Name: "https"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			DeploymentRuntimeWithAdditionalPorts(tc.args.ports)(tc.args.deployment)
			if diff := cmp.Diff(tc.want.deployment, tc.args.deployment); diff != "" {
				t.Errorf("\n%s\nDeploymentRuntimeWithAdditionalPorts(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
