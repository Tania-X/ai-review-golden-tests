"""异步会话依赖: async generator 必须用 async with 包裹才能作为上下文管理器。"""

from __future__ import annotations

from collections.abc import AsyncIterator
from contextlib import asynccontextmanager
from typing import Any


class Session:
    """数据库会话(此处用内存标记代替真实连接)。"""

    def __init__(self) -> None:
        self.closed = False

    async def execute(self, sql: str, params: dict[str, Any] | None = None) -> list[dict[str, Any]]:
        """执行参数化查询(占位符传参, 不做字符串拼接)。"""
        return [{"sql": sql, "params": params or {}}]

    def close(self) -> None:
        """释放会话。"""
        self.closed = True


async def get_session() -> AsyncIterator[Session]:
    """会话依赖(async generator: 由调用方负责关闭)。"""
    session = Session()
    try:
        yield session
    finally:
        session.close()


@asynccontextmanager
async def managed_session() -> AsyncIterator[Session]:
    """把 async generator 包成异步上下文管理器。"""
    async with get_session() as session:
        yield session


async def load_documents(names: list[str]) -> list[dict[str, Any]]:
    """在会话作用域内批量查询: 退出 async with 之前连接一定还活着。"""
    rows: list[dict[str, Any]] = []
    async with managed_session() as session:
        for name in names:
            rows.extend(await session.execute("select * from docs where name = :name", {"name": name}))
    return rows
