package payment;

import android.content.Intent;
import android.os.Bundle;

import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import java.util.List;

import cart.CartFragment;
import model.Order;
import temporary.variable.android.decentralizeddeliveryapp.R;


public class PaymentActivity extends AppCompatActivity {
    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_container);

//        if(getIntent() != null) {
//            Intent intent = getIntent();
//            List<Order> cartList = (List<Order>) intent.getSerializableExtra("cartList");
//            getSupportFragmentManager().beginTransaction().add(R.id.fragment_container, PaymentFragment.newInstance(cartList)).commit();
//        }

        getSupportFragmentManager().beginTransaction().add(R.id.fragment_container, new PaymentFragment()).commit();
    }
}
