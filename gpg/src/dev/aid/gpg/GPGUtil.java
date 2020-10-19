package dev.aid.gpg;

import com.alibaba.fastjson.JSONObject;
import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Platform;

/**
 * GPG工具类
 *
 * @author: 04637@163.com
 * @date: 2020/10/19
 */

public class GPGUtil {
    private static OpenPGP instance;

    static {
        if (Platform.isWindows()) {
            instance = Native.load("wingpg", OpenPGP.class);
        } else if (Platform.isLinux()) {
            instance = Native.load("linuxgpg", OpenPGP.class);
        }
    }

    /**
     * 通过密码加密
     *
     * @param psw    密码
     * @param toEnc  待加密文件路径
     * @param toSave 加密文件保存路径
     */
    public static Resp encryptByPsw(String psw, String toEnc, String toSave) {
        String res = instance.EncryptByPsw(psw, toEnc, toSave);
        return JSONObject.parseObject(res, Resp.class);
    }

    /**
     * 通过公钥加密
     *
     * @param pubKey 公钥文件路径
     * @param toEnc  待加密文件路径
     * @param toSave 加密文件保存路径
     */
    public static Resp encryptByPubKey(String pubKey, String toEnc, String toSave) {
        String res = instance.EncryptByPubKey(pubKey, toEnc, toSave);
        return JSONObject.parseObject(res, Resp.class);
    }

    private GPGUtil(){}
}

interface OpenPGP extends Library {
    String EncryptByPsw(String pubKey, String toEnc, String toSave);

    String EncryptByPubKey(String psw, String toEnc, String toSave);
}

