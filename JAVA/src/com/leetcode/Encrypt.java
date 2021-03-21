package com.leetcode;

public class Encrypt {
      public void en1(int[] paramArrayOfint, String paramString, int paramInt) {
    byte[] arrayOfByte2 = new byte[256];
    byte[] arrayOfByte1 = paramString.getBytes();
    int j = 0;
    int i;
    for (i = 0; i < 256; i++) {
      paramArrayOfint[i] = 256 - i;
      arrayOfByte2[i] = arrayOfByte1[i % paramInt];
    }
    i = 0;
    for (paramInt = j; paramInt < 256; paramInt++) {
      i = (paramArrayOfint[paramInt] + i + arrayOfByte2[paramInt]) % 256;
      j = paramArrayOfint[paramInt];
      paramArrayOfint[paramInt] = paramArrayOfint[i];
      paramArrayOfint[i] = j;
    }
  }

  public void en2(int[] paramArrayOfint) {
    int i = 0;
    int k = 0;
    int j = 0;
    while (i < 17) {
      k = (k + 1) % 256;
      j = ((paramArrayOfint[k] & 0xFF) + j) % 256;
      int m = paramArrayOfint[k];
      paramArrayOfint[k] = paramArrayOfint[j];
      paramArrayOfint[j] = m;
      System.out.println(paramArrayOfint[((paramArrayOfint[k] & 0xFF) + (paramArrayOfint[j] & 0xFF)) % 256]);
//      paramArrayOfbyte[i] = (byte)(paramArrayOfbyte[i] ^ paramArrayOfint[((paramArrayOfint[k] & 0xFF) + (paramArrayOfint[j] & 0xFF)) % 256]);
      i++;
    }
  }

  public boolean file(String paramString) {
          int[] arrayOfInt = new int[256];
          en1(arrayOfInt, paramString, paramString.length());
          en2(arrayOfInt);
          for (int i = 0; i < 17; i++) {
              (new int[17])[0] = 139;
              (new int[17])[1] = 210;
              (new int[17])[2] = 217;
              (new int[17])[3] = 93;
              (new int[17])[4] = 149;
              (new int[17])[5] = 255;
              (new int[17])[6] = 126;
              (new int[17])[7] = 95;
              (new int[17])[8] = 41;
              (new int[17])[9] = 86;
              (new int[17])[10] = 18;
              (new int[17])[11] = 185;
              (new int[17])[12] = 239;
              (new int[17])[13] = 236;
              (new int[17])[14] = 139;
              (new int[17])[15] = 208;
              (new int[17])[16] = 69;
            }
          return true;
      }
}
