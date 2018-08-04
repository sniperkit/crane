package crane

const (
	tmplKubernetes = ``

	tmplCrane = `{{ range .Containers }}{{ .ActualName }}:
  image: {{ .Image }}{{ if .BuildParams.Context }}
  build: {{ .BuildParams.Context }}{{ end }}{{ if .BuildParams.File }}
  dockerfile: {{ .BuildParams.File }}{{ end }}{{ if .CapAdd }}
  cap-add:{{ range .CapAdd }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.CapDrop }}
  cap-drop:{{ range .RunParams.CapDrop }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.BlkioWeightDevice }}
  blkio-weight-device:{{ range .RunParams.BlkioWeightDevice }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Cmd }}
  cmd:{{ range .RunParams.Cmd }} {{ . }}{{ end }}{{ end }}{{ if .RunParams.CgroupParent }}
  cgroup-parent: {{ .RunParams.CgroupParent }}{{ end }}{{ if .RunParams.Device }}
  devices:{{ range .RunParams.Device }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Dns }}
  dns:{{ range .RunParams.Dns }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DnsSearch }}
  dns-search:{{ range .RunParams.DnsSearch }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.EnvFile }}
  env-file:{{ range .RunParams.EnvFile }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.LxcConf }}
  lxc-conf:{{ range .RunParams.LxcConf }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Env }}
  env:{{ range .RunParams.Env }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Expose }}
  expose:{{ range .RunParams.Expose }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.AddHost }}
  add-host:{{ range .RunParams.AddHost }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Label }}
  label:{{ range .RunParams.Label }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.LabelFile }}
  label-file:{{ range .RunParams.LabelFile }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Link }}
  link:{{ range .RunParams.Link }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Ip }}
  ip: {{ .RunParams.Ip }}{{ end }}{{ if .RunParams.Ip6 }}
  ip6: {{ .RunParams.Ip6 }}{{ end }}{{ if .RunParams.Ipc }}
  ipc: {{ .RunParams.Ipc }}{{ end }}{{ if .RunParams.Isolation }}
  isolation: {{ .RunParams.Isolation }}{{ end }}{{ if .RunParams.LogDriver }}
  log-driver: {{ .RunParams.LogDriver }}{{ end }}{{ if .RunParams.LogOpt }}
  log-opt:{{ range .RunParams.LogOpt }}
  - {{ . }}{{ end }}{{ end }}{{ if ne .RunParams.Net "bridge" }}
  net: {{ .RunParams.Net }}{{ end }}{{ if .RunParams.BlkioWeight }}
  blkio-weight: {{ .RunParams.BlkioWeight }}{{ end }}{{ if .RunParams.Pid }}
  pid: {{ .RunParams.Pid }}{{ end }}{{ if .RunParams.Detach }}
  detach: {{ .RunParams.Detach }}{{ end }}{{ if .RunParams.DetachKeys }}
  detach-keys: {{ .RunParams.DetachKeys }}{{ end }}{{ if .RunParams.KernelMemory }}
  kernel-memory: {{ .RunParams.KernelMemory }}{{ end }}{{ if .RunParams.Publish }}
  publish:{{ range .RunParams.Publish }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.SecurityOpt }}
  security-opt:{{ range .RunParams.SecuirtyOpt }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Volume }}
  volume:{{ range .RunParams.Volume }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.VolumesFrom }}
  volumes-from:{{ range .RunParams.VolumesFrom }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Requires }}
	requires:{{ range .RunParams.Requires }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.CpuShares }}
  cpu-shared: {{ .RunParams.CpuShares }}{{ end }}{{ if .RunParams.CPUPeriod }}
  cpu-period: {{ .RunParams.CPUPeriod }}{{ end }}{{ if .RunParams.CPUQuota }}
  cpu-quota: {{ .RunParams.CPUQuota }}{{ end }}{{ if .RunParams.Cidfile }}
  cidfile: {{ .RunParams.Cidfile }}{{ end }}{{ if .RunParams.OomKillDisable }}
  oom-kill-disable: {{ .RunParams.OomKillDisable }}{{ end }}{{ if .RunParams.OomScoreAdj }}
  oom-score-adj: {{ .RunParams.OomScoreAdj }}{{ end }}{{ if .RunParams.ReadOnly }}
  read-only: {{ .RunParams.ReadOnly }}{{ end }}{{ if .RunParams.PublishAll }}
  publish-all: {{ .RunParams.PublishAll }}{{ end }}{{ if .RunParams.Rm }}
  rm: {{ .RunParams.Rm }}{{ end }}{{ if .RunParams.ShareSshSocket }}
  share-ssh-socket: {{ .RunParams.ShareSshSocket }}{{ end }}{{ if .RunParams.SigProxy }}
  sig-proxy: {{ .RunParams.SigProxy }}{{ end }}{{ if .RunParams.StopSignal }}
  stop-signal: {{ .RunParams.StopSignal }}{{ end }}{{ if .RunParams.Tmpfs }}
  tmpfs: {{ range .RunParams.Tmpfs }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Ulimit }}
  ulimit: {{ range .RunParams.Ulimit }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.ShmSize }}
  shm-size: {{ .RunParams.ShmSize }}{{ end }}{{ if .RunParams.Cpuset }}
  cpuset: {{ .RunParams.Cpuset }}{{ end }}{{ if .RunParams.Entrypoint }}
  entrypoint: {{ .RunParams.Entrypoint }}{{ end }}{{ if .RunParams.User }}
  user: {{ .RunParams.User }}{{ end }}{{ if .RunParams.Uts }}
  uts: {{ .RunParams.Uts }}{{ end }}{{ if .RunParams.Workdir }}
  workdir: {{ .RunParams.Workdir }}{{ end }}{{ if .RunParams.HealthCmd }}
  health-cmd: {{ .RunParams.HealthCmd }}{{ end }}{{ if .RunParams.HealthInterval }}
  health-interval: {{ .RunParams.HealthInterval }}{{ end }}{{ if .RunParams.HealthRetries }}
  health-retries: {{ .RunParams.HealthRetries }}{{ end }}{{ if .RunParams.HealthTimeout }}
  health-timeouts: {{ .RunParams.HealthTimeout }}{{ end }}{{ if .RunParams.Hostname }}
  hostname: {{ .RunParams.Hostname }}{{ end }}{{ if .RunParams.MacAddress }}
  mac-address: {{ .RunParams.MacAddress }}{{ end }}{{ if .RunParams.Memory }}
  memory-limit: {{ .RunParams.Memory }}{{ end }}{{ if .RunParams.MemorySwap }}
  memswap-limit: {{ .RunParams.MemorySwap }}{{ end }}{{ if .RunParams.Privileged }}
  privileged: {{ .RunParams.Privileged }}{{ end }}{{ if .RunParams.Restart }}
  restart: {{ .RunParams.Restart }}{{ end }}{{ if .RunParams.ReadOnly }}
  read_only: {{ .RunParams.ReadOnly }}{{ end }}{{ if .RunParams.Init }}
  init: {{ .RunParams.Init }}{{ end }}{{ if .RunParams.Interactive }}
  interactive: {{ .RunParams.Interactive }}{{ end }}{{ if .RunParams.Net }}
  net: {{ .RunParams.Net }}{{ end }}{{ if .RunParams.VolumeDriver }}
  volume-driver: {{ .RunParams.VolumeDriver }}{{ end }}{{ if .RunParams.Tty }}
  tty: {{ .RunParams.Tty }}{{ end }}{{ if .RunParams.GroupAdd }}
  group-add:{{ range .RunParams.GroupAdd }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceReadBps }}
  device-read-bps:{{ range .RunParams.DeviceReadBps }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceReadIops }}
  device-read-iops:{{ range .RunParams.DeviceReadIops }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceWriteBps }}
  device-write-bp:{{ range .RunParams.DeviceWriteBps }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceWriteIops }}
  device-write-iops:{{ range .RunParams.DeviceWriteIops }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceWriteIops }}
  device-write-iops:{{ range .RunParams.DeviceWriteIops }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Networks }}
  networks:{{ range .RunParams.Networks }}
  - {{ . }}{{ end }}{{ end }}
{{ end }}`

	tmplDockerCompose = `{{ range .Containers }}{{ .ActualName }}:
  image: {{ .Image }}{{ if .BuildParams.Context }}
  build: {{ .BuildParams.Context }}{{ end }}{{ if .BuildParams.File }}
  dockerfile: {{ .BuildParams.File }}{{ end }}{{ if .RunParams.CapAdd }}
  cap_add:{{ range .RunParams.CapAdd }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.CapDrop }}
  cap_drop:{{ range .RunParams.CapDrop }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.BlkioWeightDevice }}
  blkio-weight-device:{{ range .RunParams.BlkioWeightDevice }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Cmd }}
  command:{{ range .RunParams.Cmd }} {{ . }}{{ end }}{{ end }}{{ if .RunParams.CgroupParent }}
  cgroup_parent: {{ .RunParams.CgroupParent }}{{ end }}{{ if .RunParams.Device }}
  device:{{ range .RunParams.Device }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Dns }}
  dns:{{ range .RunParams.Dns }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DnsSearch }}
  dns_search:{{ range .RunParams.DnsSearch }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.EnvFile }}
  env_file:{{ range .RunParams.EnvFile }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.LxcConf }}
  lxc-conf:{{ range .RunParams.LxcConf }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Env }}
  environment:{{ range .RunParams.Env }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Expose }}
  expose:{{ range .RunParams.Expose }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.AddHost }}
  extra_hosts:{{ range .RunParams.AddHost }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Label }}
  labels:{{ range .RunParams.Label }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.LabelFile }}
  label-file:{{ range .RunParams.LabelFile }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Link }}
  links:{{ range .RunParams.Link }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Ip }}
  ip: {{ .RunParams.Ip }}{{ end }}{{ if .RunParams.Ip6 }}
  ip6: {{ .RunParams.Ip6 }}{{ end }}{{ if .RunParams.Ipc }}
  ipc: {{ .RunParams.Ipc }}{{ end }}{{ if .RunParams.Isolation }}
  isolation: {{ .RunParams.Isolation }}{{ end }}{{ if .RunParams.LogDriver }}
  log-driver: {{ .RunParams.LogDriver }}{{ end }}{{ if .RunParams.LogOpt }}
  log-opt:{{ range .RunParams.LogOpt }}
  - {{ . }}{{ end }}{{ end }}{{ if ne .RunParams.Net "bridge" }}
  net: {{ .RunParams.Net }}{{ end }}{{ if .RunParams.BlkioWeight }}
  blkio-weight: {{ .RunParams.BlkioWeight }}{{ end }}{{ if .RunParams.Pid }}
  pid: {{ .RunParams.Pid }}{{ end }}{{ if .RunParams.Detach }}
  detach: {{ .RunParams.Detach }}{{ end }}{{ if .RunParams.DetachKeys }}
  detach-keys: {{ .RunParams.DetachKeys }}{{ end }}{{ if .RunParams.KernelMemory }}
  kernel-memory: {{ .RunParams.KernelMemory }}{{ end }}{{ if .RunParams.Publish }}
  publish:{{ range .RunParams.Publish }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.SecurityOpt }}
  security_opt:{{ range .RunParams.SecuirtyOpt }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Volume }}
  volumes:{{ range .RunParams.Volume }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.VolumesFrom }}
  volumes-from:{{ range .RunParams.VolumesFrom }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Requires }}
  depends_on:{{ range .RunParams.Requires }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.CpuShares }}
  cpu-shared: {{ .RunParams.CpuShares }}{{ end }}{{ if .RunParams.CPUPeriod }}
  cpu-period: {{ .RunParams.CPUPeriod }}{{ end }}{{ if .RunParams.CPUQuota }}
  cpu-quota: {{ .RunParams.CPUQuota }}{{ end }}{{ if .RunParams.Cidfile }}
  cidfile: {{ .RunParams.Cidfile }}{{ end }}{{ if .RunParams.OomKillDisable }}
  oom-kill-disable: {{ .RunParams.OomKillDisable }}{{ end }}{{ if .RunParams.OomScoreAdj }}
  oom-score-adj: {{ .RunParams.OomScoreAdj }}{{ end }}{{ if .RunParams.ReadOnly }}
  read_only: {{ .RunParams.ReadOnly }}{{ end }}{{ if .RunParams.PublishAll }}
  publish-all: {{ .RunParams.PublishAll }}{{ end }}{{ if .RunParams.Rm }}
  rm: {{ .RunParams.Rm }}{{ end }}{{ if .RunParams.ShareSshSocket }}
  share-ssh-socket: {{ .RunParams.ShareSshSocket }}{{ end }}{{ if .RunParams.SigProxy }}
  sig-proxy: {{ .RunParams.SigProxy }}{{ end }}{{ if .RunParams.StopSignal }}
  stop_signal: {{ .RunParams.StopSignal }}{{ end }}{{ if .RunParams.Tmpfs }}
  tmpfs: {{ range .RunParams.Tmpfs }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Ulimit }}
  ulimit: {{ range .RunParams.Ulimit }}
  - {{ . }}{{ end }}{{ end }}{{ if .RunParams.ShmSize }}
  shm_size: {{ .RunParams.ShmSize }}{{ end }}{{ if .RunParams.Cpuset }}
  cpuset: {{ .RunParams.Cpuset }}{{ end }}{{ if .RunParams.Entrypoint }}
  entrypoint: {{ .RunParams.Entrypoint }}{{ end }}{{ if .RunParams.User }}
  user: {{ .RunParams.User }}{{ end }}{{ if .RunParams.Uts }}
  uts: {{ .RunParams.Uts }}{{ end }}{{ if .RunParams.Workdir }}
  working_dir: {{ .RunParams.Workdir }}{{ end }}{{ if .RunParams.HealthCmd }}
  health-cmd: {{ .RunParams.HealthCmd }}{{ end }}{{ if .RunParams.HealthInterval }}
  health-interval: {{ .RunParams.HealthInterval }}{{ end }}{{ if .RunParams.HealthRetries }}
  health-retries: {{ .RunParams.HealthRetries }}{{ end }}{{ if .RunParams.HealthTimeout }}
  health-timeouts: {{ .RunParams.HealthTimeout }}{{ end }}{{ if .RunParams.Hostname }}
  hostname: {{ .RunParams.Hostname }}{{ end }}{{ if .RunParams.MacAddress }}
  mac_address: {{ .RunParams.MacAddress }}{{ end }}{{ if .RunParams.Memory }}
  memory-limit: {{ .RunParams.Memory }}{{ end }}{{ if .RunParams.MemorySwap }}
  memswap-limit: {{ .RunParams.MemorySwap }}{{ end }}{{ if .RunParams.Privileged }}
  privileged: {{ .RunParams.Privileged }}{{ end }}{{ if .RunParams.Restart }}
  restart: {{ .RunParams.Restart }}{{ end }}{{ if .RunParams.ReadOnly }}
  read_only: {{ .RunParams.ReadOnly }}{{ end }}{{ if .RunParams.Init }}
  init: {{ .RunParams.Init }}{{ end }}{{ if .RunParams.Interactive }}
  stdin_open: {{ .RunParams.Interactive }}{{ end }}{{ if .RunParams.Net }}
  network_mode: {{ .RunParams.Net }}{{ end }}{{ if .RunParams.Tty }}
  tty: {{ .RunParams.Tty }}{{ end }}{{ if .RunParams.VolumeDriver }}
  volume-driver: {{ .RunParams.VolumeDriver }}{{ end }}{{ if .RunParams.GroupAdd }}
  group_add:{{ range .RunParams.GroupAdd }}
    - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceReadBps }}
  device-read-bps:{{ range .RunParams.DeviceReadBps }}
    - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceReadIops }}
  device-read-iops:{{ range .RunParams.DeviceReadIops }}
    - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceWriteBps }}
  device-write-bp:{{ range .RunParams.DeviceWriteBps }}
    - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceWriteIops }}
  device-write-iops:{{ range .RunParams.DeviceWriteIops }}
    - {{ . }}{{ end }}{{ end }}{{ if .RunParams.DeviceWriteIops }}
  device-write-iops:{{ range .RunParams.DeviceWriteIops }}
    - {{ . }}{{ end }}{{ end }}{{ if .RunParams.Networks }}
  networks:{{ range .RunParams.Networks }}
    - {{ . }}{{ end }}{{ end }}
{{ end }}`

	tmplSystemD = `[Unit]
Description={{ .Name }} container
Requires=docker.service{{ range .Dependencies.All }} {{ . }}.service{{ end }}
After=docker.service{{ range .Dependencies.All }} {{ . }}.service{{ end }}

[Service]
EnvironmentFile=/path/to/params.env
Restart=on-failure
RestartSec=10
ExecStartPre=/bin/sleep 2
# Be sure that every 'detach' attribute in crane.yml is unset
ExecStart=/path/to/crane --config /path/to/crane.yml run {{ .Name }}
ExecStop=/path/to/crane --config /path/to/crane.yml stop {{ .Name }}

[Install]
WantedBy=multi-user.target`

	tmplUpStart = `description "{{ .Name }} container"
author "Me"
start on filesystem and started docker
stop on runlevel [!2345]
respawn
script
 /path/to/crane --config /path/to/crane.yml run {{ .Name }}
end script`

	tmplDot = `digraph {
{{ range $container := .Containers }}  "{{ .Name }}" [style=bold]
{{ range .Dependencies.Link }}  "{{ $container.Name }}"->"{{ . }}"
{{ end }}{{ range .Dependencies.VolumesFrom }}  "{{ $container.Name }}"->"{{ . }}" [style=dashed]
{{ end }}{{ if .Dependencies.Net }}  "{{ $container.Name }}"->"{{ . }}" [style=dotted]
{{ end }}{{ end }}}`

	tmplServiceUnit = `
[Unit]
Description=Start, stop and restart {{name}} container
Requires=docker.service {{#links}}{{.}} {{/links}}
After=docker.service {{#links}}{{.}} {{/links}}

[Service]
Restart=always

ExecStartPre=-/usr/bin/docker rm -f %n

ExecStart=/usr/bin/docker run \
    {{#environment}}-e {{{.}}} {{/environment}}\
    {{#links}}--link {{.}}:{{.}} {{/links}} \
    {{#ports}}-p {{.}} {{/ports}} \
    {{#env_file}}--env-file={{{.}}} {{/env_file}} \
    --rm --name %n \
    {{#image}}{{{.}}}{{/image}} {{#command}}{{{.}}}{{/command}}

ExecStop=-/usr/bin/docker stop %n

[Install]
WantedBy=multi-user.target
`
	tmplOneshotUnit = `
[Unit]
Description=Start a oneshot for {{name}} container

[Service]
Type=oneshot

ExecStart=/usr/bin/docker exec \
    {{#exec}}{{{.}}}{{/exec}} {{#command}}{{{.}}}{{/command}}
`
)
