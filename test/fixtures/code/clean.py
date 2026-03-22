"""Clean generated code that passes security scans."""
import json


def process_data(data: dict) -> str:
    return json.dumps(data, indent=2)


def validate_input(value: str) -> bool:
    return bool(value and len(value) < 1000)
