package dev.aid.gpg;

public class Resp {
    private int code;
    private String msg;
    private String data;

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public void setData(String data) {
        this.data = data;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getMsg() {
        return msg;
    }

    public int getCode() {
        return code;
    }

    public String getData() {
        return data;
    }
}
