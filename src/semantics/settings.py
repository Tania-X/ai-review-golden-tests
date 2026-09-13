"""配置读取:`x or default` 里的默认值来自配置对象, 不是硬编码常量。"""

from __future__ import annotations

import os
from dataclasses import dataclass


@dataclass(frozen=True, slots=True)
class Settings:
    """运行期配置(从环境变量读取)。"""

    base_url: str
    timeout: float


def _env_float(name: str, default: float) -> float:
    """读取浮点配置: 未设置或写了非法值时回退到默认值。"""
    try:
        return float(os.environ.get(name, ""))
    except ValueError:
        return default


def load_settings() -> Settings:
    """加载配置: 缺省值集中在配置层。"""
    return Settings(
        base_url=os.environ.get("SERVICE_BASE_URL", "http://localhost:8080"),
        timeout=_env_float("SERVICE_TIMEOUT", 30.0),
    )


settings = load_settings()


def client_for(base_url: str | None = None, *, timeout: float | None = None) -> dict[str, object]:
    """构造客户端: 未显式传入时回退到配置值(而非内置常量)。"""
    resolved_base = base_url or settings.base_url
    resolved_timeout = timeout or settings.timeout
    return {"base_url": resolved_base, "timeout": resolved_timeout}
