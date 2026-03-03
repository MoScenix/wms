# 快速开始

## 1. 启动基础依赖

```bash
docker compose up -d consul mysql redis
```

## 2. 启动后端服务

分别启动：

- `app/user`
- `app/wms`
- `app/bff`

示例：

```bash
cd app/user && go run .
cd app/wms && go run .
cd app/bff && go run .
```

## 3. 启动前端

```bash
cd app/frontend
pnpm install
pnpm dev
```

默认本地访问：`http://localhost:5173`。
