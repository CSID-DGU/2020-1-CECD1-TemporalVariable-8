package cart;

import android.content.Intent;
import android.os.Bundle;

import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import java.util.List;

import model.Order;
import model.Restaurant;
import temporary.variable.android.decentralizeddeliveryapp.R;

public class CartActivity extends AppCompatActivity {
    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_container);

        if(getIntent() != null) {
            Intent intent = getIntent();
//            List<Order> cartList = (List<Order>) intent.getSerializableExtra("cartList");
//            Restaurant restaurant = (Restaurant) intent.getSerializableExtra("restaurant");
            List<String> restaurant = (List<String>) intent.getSerializableExtra("restaurant");
//            getSupportFragmentManager().beginTransaction().add(R.id.fragment_container, CartFragment.newInstance(restaurant, cartList)).commit();
            getSupportFragmentManager().beginTransaction().add(R.id.fragment_container, CartFragment.newInstance(restaurant)).commit();
        }
    }
}
