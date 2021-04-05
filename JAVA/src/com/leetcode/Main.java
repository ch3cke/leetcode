package com.leetcode;

import java.nio.charset.StandardCharsets;

public class Main {
  public static void main(String[] args) {
        int v11 = 0x1F;
        int v9 = 2;
        int v2 = 1;
//        String v6 = this.val$editview.getText().toString();
//        if(v6.length() != 0x20 || v6.charAt(v11) != 97 || v6.charAt(1) != 98 || v6.charAt(0) + v6.charAt(v9) - 0x30 != 56) {
//            v2 = 0;
//        }
                    char[] v5 = "dd2940c04462b4dd7c450528835cca15".toCharArray();
                    v5[v9] = ((char)(v5[v9] + v5[3] - 50));
                    v5[4] = ((char)(v5[v9] + v5[5] - 0x30));
                    v5[30] = ((char)(v5[v11] + v5[9] - 0x30));
                    v5[14] = ((char)(v5[27] + v5[28] - 97));
                    int v4;
                    for(v4 = 0; v4 < 16; ++v4) {
                        char v0 = v5[0x1F - v4];
                        v5[0x1F - v4] = v5[v4];
                        v5[v4] = v0;
                    }

        System.out.println(v5);
    }
}
