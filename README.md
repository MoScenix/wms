# WMS v1

WMS v1 is implemented with the same stack and layering as `ai-code`:

- `idl -> rpc_gen -> app/*`
- Kitex (RPC)
- Hertz (BFF)
- Consul + MySQL + Redis

## Scope

- Inbound orders
- Outbound orders
- Inventory (warehouse + location + lot + expiry)
- Inventory transaction log

## USER Reuse

- `idl/user.proto` is reused as-is.
- BFF keeps `userservice.Client` and forwards identity via metainfo.
- Session based auth is wired in `app/bff/middleware`.

## Layout

- `idl/`: API contracts (`user.proto`, `wms.proto`, `bff/wms_bff.proto`)
- `rpc_gen/`: generated Kitex clients/stubs
- `app/wms/`: WMS RPC server
- `app/bff/`: WMS BFF HTTP gateway
- `common/`: shared observability and server suites

## Notes

- `app/wms/biz/store` currently provides an in-memory transactional implementation for fast bootstrap and interface verification.
- DAL bootstrap for MySQL/Redis is present and can be switched to persistent model implementations incrementally.
# wms
