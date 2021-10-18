package temporary.variable.android.decentralizeddeliveryapp;

import android.content.Intent;
import android.os.Bundle;
import android.os.Handler;

import androidx.appcompat.app.AppCompatActivity;


public class SplashActivity extends AppCompatActivity {
    private static final int SPLASH_TIME_OUT = 2000;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_container);

        new Handler().postDelayed(() -> {
            Intent i = new Intent(getApplicationContext(), MainActivity.class);
            startActivity(i);

            finish();
        }, SPLASH_TIME_OUT);
    }
}
