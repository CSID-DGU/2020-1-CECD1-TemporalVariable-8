package cart;

import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.os.Parcelable;
import android.util.DisplayMetrics;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.ScrollView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.DividerItemDecoration;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import java.io.Serializable;
import java.util.List;

import model.Order;
import model.Restaurant;
import payment.PaymentActivity;
//import restaurant_detail.RestaurantDetailFragment;
import temporary.variable.android.decentralizeddeliveryapp.R;
import util.Util;

//import static restaurant_detail.RestaurantDetailFragment.mCartList;
import static restaurant_detail.RestaurantDetailActivity.mCartList;

public class CartFragment extends Fragment {
    private TextView mRestaurantName;
    private TextView mTotalPrice;
    private Button mBtnOrder;

    //    private Restaurant mRestaurant;
    private List<String> mRestaurant;

//    private List<Order> mCartList;

    private RecyclerView mRecyclerView;
    private CartRecyclerAdapter mCartRecyclerAdapter;
    private CartActivity mContext;

    public static CartFragment newInstance(Restaurant restaurant, List<Order> list) {
        CartFragment fragment = new CartFragment();
        Bundle args = new Bundle();
        args.putSerializable("restaurant", (Serializable) restaurant);
        args.putSerializable("cartList", (Serializable) list);
        fragment.setArguments(args);
        return fragment;
    }

//    public static CartFragment newInstance(Restaurant restaurant) {
//        CartFragment fragment = new CartFragment();
//        Bundle args = new Bundle();
//        args.putSerializable("restaurant", (Serializable) restaurant);
//        fragment.setArguments(args);
//        return fragment;
//    }

    public static CartFragment newInstance(List<String> restaurant) {
        CartFragment fragment = new CartFragment();
        Bundle args = new Bundle();
        args.putSerializable("restaurant", (Serializable) restaurant);
        fragment.setArguments(args);
        return fragment;
    }

    public CartFragment() {
    }

    @Override
    public void onAttach(Context context) {
        super.onAttach(context);
        mContext = (CartActivity) context;
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = null;


        if (getArguments() != null) {
//            mRestaurant = (Restaurant) getArguments().getSerializable("restaurant");
            mRestaurant = (List<String>) getArguments().getSerializable("restaurant");

//            mCartList = (List<Order>) getArguments().getSerializable("cartList");

            if (mCartList.size() > 0) {
                v = inflater.inflate(R.layout.fragment_cart, container, false);

                mRestaurantName = v.findViewById(R.id.textView_cart_restaurant_name);
                mTotalPrice = v.findViewById(R.id.textView_cart_total_price);
                mBtnOrder = v.findViewById(R.id.btn_order);
                mRecyclerView = v.findViewById(R.id.cart_recyclerView);

                mTotalPrice.setText(Util.numberToCommaFormat(calculateTotalPrice(mCartList)));
                if(mRestaurant.size() > 1)
                    mRestaurantName.setText(mRestaurant.get(0) + " " + mRestaurant.get(1));
                else
                    mRestaurantName.setText(mRestaurant.get(0));

//                mRestaurantName.setText(mRestaurant.getName() + " " + mRestaurant.getLocation());

                mCartRecyclerAdapter = new CartRecyclerAdapter();
                mCartRecyclerAdapter.setCartList(mCartList);
                mRecyclerView.setLayoutManager(new LinearLayoutManager(getContext()));
                mRecyclerView.addItemDecoration(new DividerItemDecoration(getContext(), LinearLayoutManager.VERTICAL));
                mRecyclerView.setAdapter(mCartRecyclerAdapter);

                mBtnOrder.setOnClickListener(view -> {
                    Intent intent = new Intent(getContext(), PaymentActivity.class);
//                    intent.putExtra("cartList", (Serializable) mCartList);
                    startActivity(intent);
                });
            } else {
                v = inflater.inflate(R.layout.fragment_empty, container, false);
            }
        }

        return v;
    }

    private int calculateTotalPrice(List<Order> list) {
        int price = 0;
        for (Order order : list) price += order.getOrderPrice();

        return price;
    }

    class CartRecyclerAdapter extends RecyclerView.Adapter<CartViewHolder> {
//        List<Order> cartList;

        public CartRecyclerAdapter() {
//            cartList = list;
        }

        @NonNull
        @Override
        public CartViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
            View v = LayoutInflater.from(parent.getContext()).inflate(R.layout.list_item_cart, parent, false);
            return new CartViewHolder(v);
        }

        @Override
        public void onBindViewHolder(@NonNull CartViewHolder holder, int i) {
            final Order order = mCartList.get(i);

            holder.menuName.setText(order.getMenu().getName());
            holder.count.setText("수량 : " + order.getCount());
            holder.price.setText(Util.numberToCommaFormat(order.getOrderPrice()));
            holder.btnDelete.setOnClickListener(view -> {
                mCartList.remove(i);
                mTotalPrice.setText(Util.numberToCommaFormat(calculateTotalPrice(mCartList)));
                setCartList(mCartList);
                Toast.makeText(getContext(), "삭제되었습니다.", Toast.LENGTH_SHORT).show();
            });
        }

        @Override
        public int getItemCount() {
            return mCartList.size();
        }

        public void setCartList(List<Order> list) {
            mCartList = list;
            notifyDataSetChanged();
        }
    }

    class CartViewHolder extends RecyclerView.ViewHolder {
        TextView menuName;
        TextView count;
        TextView price;
        Button btnDelete;

        public CartViewHolder(@NonNull View itemView) {
            super(itemView);
            menuName = itemView.findViewById(R.id.textView_cart_item_menu_name);
            count = itemView.findViewById(R.id.textView_cart_item_menu_count);
            price = itemView.findViewById(R.id.textView_cart_item_menu_price);
            btnDelete = itemView.findViewById(R.id.btn_cart_item_menu_delete);
        }
    }
}
