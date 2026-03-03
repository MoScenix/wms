/**
 * 环境变量配置说明
 *
 * 在项目根目录创建 .env.local 或 .env.development 文件，并添加以下配置：
 *
 * # 应用部署域名
 * VITE_DEPLOY_DOMAIN=http://localhost
 *
 * # API 基础地址（同域反向代理推荐 /api）
 * VITE_API_BASE_URL=/api
 *
 * # 本地开发时，Vite 将 /api 代理到该地址（默认 http://localhost:8080）
 * VITE_API_PROXY_TARGET=http://localhost:8080
 *
 * 生产环境可以创建 .env.production 文件：
 *
 * # 生产环境配置示例
 * VITE_DEPLOY_DOMAIN=https://your-domain.com
 * VITE_API_BASE_URL=/api
 */

export {}
