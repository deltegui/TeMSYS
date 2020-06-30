package com.deltegui.temsys;

public class UseCaseException extends RuntimeException {
    private final int code;
    private final String reason;
    private final String fix;

    public UseCaseException(int code, String message, String fix) {
        super(message);
        this.code = code;
        this.reason = message;
        this.fix = fix;
    }

    public int getCode() {
        return code;
    }

    public String getReason() {
        return reason;
    }

    public String getFix() {
        return fix;
    }
}
