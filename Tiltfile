load('ext://pack', 'pack')

pack('pods', builder='paketobuildpacks/builder:base', env_vars=['BP_GO_TARGETS=./cmd/pods'])
pack('webhook', builder='paketobuildpacks/builder:base', env_vars=['BP_GO_TARGETS=./cmd/webhook'])

k8s_yaml('deployments/tilt.yaml')

k8s_resource('pods', port_forwards=8000)
k8s_resource('webhook', port_forwards=8080)
