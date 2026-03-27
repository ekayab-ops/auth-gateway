import logging
import base64
import hmac
import hashlib
import json

def generate_jwt_payload(user_id, expires_at, secret_key):
    """
    Generate a JWT payload with the provided user ID and expiration time.

    :param user_id: The ID of the user.
    :param expires_at: The expiration time of the JWT.
    :param secret_key: The secret key used for signing the JWT.
    :return: A dictionary containing the JWT payload.
    """
    payload = {
        "user_id": user_id,
        "exp": expires_at,
    }
    return payload

def sign_jwt(payload, secret_key):
    """
    Sign a JWT payload with the provided secret key.

    :param payload: The JWT payload to be signed.
    :param secret_key: The secret key used for signing the JWT.
    :return: A string representing the signed JWT.
    """
    encoded_payload = json.dumps(payload).encode("utf-8")
    signature = base64.urlsafe_b64encode(hmac.new(secret_key.encode("utf-8"), encoded_payload, hashlib.sha256).digest())
    return signature

def verify_jwt(token, secret_key):
    """
    Verify a JWT token with the provided secret key.

    :param token: The JWT token to be verified.
    :param secret_key: The secret key used for verifying the JWT.
    :return: A dictionary containing the verified JWT payload, or None if the token is invalid.
    """
    try:
        payload = json.loads(base64.urlsafe_b64decode(token).decode("utf-8"))
        encoded_payload = json.dumps(payload).encode("utf-8")
        signature = hmac.new(secret_key.encode("utf-8"), encoded_payload, hashlib.sha256).digest()
        if base64.urlsafe_b64encode(signature) != token:
            raise ValueError("Invalid signature")
        return payload
    except (ValueError, json.JSONDecodeError):
        return None