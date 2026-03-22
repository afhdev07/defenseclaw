"""
TEST FIXTURE — intentionally contains a hardcoded secret.
DO NOT use as a template. This file exists to test CodeGuard detection.
"""
import requests

# INTENTIONAL: hardcoded credential for scanner testing
_TEST_API_KEY = "sk-test-1234567890abcdef1234567890abcdef"  # noqa: S105


def call_api() -> dict:
    headers = {"Authorization": f"Bearer {_TEST_API_KEY}"}
    response = requests.get("https://api.example.com/data", headers=headers, timeout=10)
    return response.json()
