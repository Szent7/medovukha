<div align="center">
<img src="./frontend/src/assets/logo_small.svg" width="256" alt="" />
  <h1> Medovukha 🍯</h1>
  <p> Open-source web interface for Docker </p>
</div>

# 🚧 This repository is archived!

> The original **medovukha** project has been split into two new repositories because
> some Docker client functionality have become legacy and adding integrations is
> difficult when the app runs inside a Docker container instead of on the host machine.

## New locations

| Repository | Purpose | Link |
|------------|---------|------|
| **medovukha-core** | The core logic and helper utilities | [github.com/Szent7/medovukha-core](https://github.com/Szent7/medovukha-core) |
| **medovukha-web**  | The web frontend | [github.com/Szent7/medovukha-web](https://github.com/Szent7/medovukha-web) |

## Why the change?

* **Docker client updates** – The Docker client used in the original repo contains legacy components that force us to execute commands directly on the host machine (`exec.CommandContext`).
* **Integration difficulty** – Adding new integrations will be far easier if the application runs as a Linux daemon rather than inside a container.

## 🔧 Install (from source)
 
### 🐳 Docker
```bash
git clone https://github.com/Szent7/medovukha.git
cd medovukha
#choose standalone or compose
docker build --tag medovukha .
docker run -d --restart=always -p 10015:10015 -v /var/run/docker.sock:/var/run/docker.sock --name medovukha medovukha:latest
#OR
docker compose up -d
