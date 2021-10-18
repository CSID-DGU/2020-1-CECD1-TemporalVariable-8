package util;

import android.app.Activity;
import android.content.Context;
import android.util.DisplayMetrics;
import android.util.Log;

import java.text.DecimalFormat;

public class Util {
    public static String numberToCommaFormat(int num) {
        DecimalFormat df = new DecimalFormat("#,###");
        return df.format(num);
    }

    public static String getDeviceSize(Activity activity) {
        DisplayMetrics displayMetrics = new DisplayMetrics();
        activity.getWindowManager().getDefaultDisplay().getMetrics(displayMetrics);
        int width = displayMetrics.widthPixels;// 가로
        int height = displayMetrics.heightPixels;// 세로
        Log.d("MYTAG", "디바이스 가로 : " + width);
        Log.d("MYTAG", "디바이스 세로 : " + height);

        return width + " " + height;
    }
}
