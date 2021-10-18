package restaurant_detail;

import android.content.Intent;
import android.graphics.Color;
import android.os.Bundle;
import android.util.Log;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.viewpager2.adapter.FragmentStateAdapter;
import androidx.viewpager2.widget.ViewPager2;

import com.google.android.material.appbar.CollapsingToolbarLayout;
import com.google.android.material.bottomsheet.BottomSheetBehavior;
import com.google.android.material.floatingactionbutton.FloatingActionButton;
import com.google.android.material.tabs.TabLayout;
import com.google.android.material.tabs.TabLayoutMediator;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

import cart.CartActivity;
import model.Menu;
import model.Order;
import model.Restaurant;
import temporary.variable.android.decentralizeddeliveryapp.R;
import util.Util;

public class RestaurantDetailActivity extends AppCompatActivity {
    private CollapsingToolbarLayout mCollapsingToolbarLayout;
    private Toolbar mToolbar;
    private FloatingActionButton mCartFab;

    private TabLayout mTabLayout;
    private ViewPager2 mViewPager2;
    private FragmentStateAdapter mPagerAdapter;
    private int pageCnt = 2;

    private TextView mSubtitle;

    public static BottomSheetBehavior mBottomSheetBehavior;
    public static LinearLayout mLinearLayout;

    public static ImageView mMinus;
    public static ImageView mPlus;

    public static TextView mPrice;
    public static TextView mCount;
    public static TextView mMenuName;
    public static TextView mCartCount;

    public static Button mButtonAddCart;

    public static int cnt = 1;
    public static int price = 0;

    public static List<Order> mCartList;
    public static Menu mMenu;

    @Override
    protected void onResume() {
        super.onResume();

        if (mCartList.size() > 0)
            mCartCount.setText(String.valueOf(calculateTotalCnt()));

        else
            mCartCount.setText("");
    }

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_restaurant_detail);
        mCartList = new ArrayList<>();

        List<String> restaurant = (List<String>) getIntent().getSerializableExtra("restaurant");

        mCartCount = findViewById(R.id.textView_cart_count);
        mMenuName = findViewById(R.id.textView_menu_name);
        mPrice = findViewById(R.id.textView_price);
        mCount = findViewById(R.id.textView_cnt);
        mMinus = findViewById(R.id.imageView_cnt_minus);
        mPlus = findViewById(R.id.imageView_cnt_plus);
        mButtonAddCart = findViewById(R.id.btn_add_cart);

        mMinus.setOnClickListener(view -> {
            if (cnt >= 2) {
                cnt--;
                price -= mMenu.getPrice();
            }

            mPrice.setText(Util.numberToCommaFormat(price) + "원");
            mCount.setText(String.valueOf(cnt));
        });

        mPlus.setOnClickListener(view -> {
            cnt++;
            price += mMenu.getPrice();

            mPrice.setText(Util.numberToCommaFormat(price) + "원");
            mCount.setText(String.valueOf(cnt));
        });

        mButtonAddCart.setOnClickListener(view -> {
            Menu menu = new Menu(mMenu.getId(), mMenu.getName(), mMenu.getPrice(), mMenu.getDescription(), mMenu.getImage());

            int value = isExistMenu(menu.getName());

            if (value >= 0) { //이미 카트에 존재하는 메뉴 추가
                Log.d("MYTAG", "cart size : " + mCartList.size());
                Order o = mCartList.get(value);
                o.setCount(o.getCount() + cnt);
                o.setOrderPrice(o.getOrderPrice() + price);
                mCartList.set(value, o);

                Log.d("MYTAG", "cart size : " + mCartList.size());

            } else mCartList.add(new Order(menu, cnt, price));

            mBottomSheetBehavior.setState(BottomSheetBehavior.STATE_HIDDEN);
            mCartCount.setText(String.valueOf(calculateTotalCnt()));
            Toast.makeText(getApplicationContext(), "장바구니에 추가되었습니다.", Toast.LENGTH_SHORT).show();
        });

        mLinearLayout = findViewById(R.id.bottom_sheet_layout);
        mBottomSheetBehavior = BottomSheetBehavior.from(mLinearLayout);

        mCartFab = findViewById(R.id.fab_cart);
        mCartFab.setOnClickListener(view -> {
            Intent intent = new Intent(getApplicationContext(), CartActivity.class);
//            intent.putExtra("cartList", (Serializable) mCartList);

//            intent.putExtra("restaurant", (Serializable) restaurant);
            intent.putExtra("restaurant", (Serializable) restaurant);

            startActivity(intent);
        });

        mCollapsingToolbarLayout = findViewById(R.id.toolbar_layout);
        mCollapsingToolbarLayout.setTitle(restaurant.get(0));
        mCollapsingToolbarLayout.setBackgroundResource(R.drawable.china_1);

//        mCollapsingToolbarLayout.setTitle(restaurant.getName());
//        mCollapsingToolbarLayout.setBackgroundResource(restaurant.getImages().get(0));
        mCollapsingToolbarLayout.setExpandedTitleColor(Color.BLACK);

        mSubtitle = findViewById(R.id.textView_toolbar_subtitle);
//        mSubtitle.setText(restaurant.getLocation());

        if(restaurant.size() > 1)
            mSubtitle.setText(restaurant.get(1));
        else
            mSubtitle.setText("본점");

        mToolbar = findViewById(R.id.toolbar);

//        mToolbar.setTitle(restaurant.getName() + " " + restaurant.getLocation());
        if(restaurant.size() > 1)
            mToolbar.setTitle(restaurant.get(0) + " " + restaurant.get(1));
        else
            mToolbar.setTitle(restaurant.get(0));

        mTabLayout = findViewById(R.id.tabLayout);
        mTabLayout.addOnTabSelectedListener(new TabLayout.OnTabSelectedListener() {
            @Override
            public void onTabSelected(TabLayout.Tab tab) {
                Log.d("MYTAG", "position : " + tab.getPosition());
            }

            @Override
            public void onTabUnselected(TabLayout.Tab tab) {

            }

            @Override
            public void onTabReselected(TabLayout.Tab tab) {

            }
        });

        mViewPager2 = findViewById(R.id.view_pager);
//        mPagerAdapter = new MyViewPagerAdapter(this, pageCnt, restaurant);

        if(restaurant.size() > 1)
            mPagerAdapter = new MyViewPagerAdapter(this, pageCnt, restaurant.get(0) + " " + restaurant.get(1));
        else
            mPagerAdapter = new MyViewPagerAdapter(this, pageCnt, restaurant.get(0));
        mViewPager2.setAdapter(mPagerAdapter);

        mViewPager2.setOrientation(ViewPager2.ORIENTATION_HORIZONTAL);
        mViewPager2.setCurrentItem(0);

        new TabLayoutMediator(mTabLayout, mViewPager2, true, (tab, position) -> {
            switch (position) {
                case 0:
                    tab.setText("메뉴");
                    break;
                case 1:
                    tab.setText("리뷰");
                    break;
            }
        }).attach();

//        getSupportFragmentManager().beginTransaction().add(R.id.fragment_container, RestaurantDetailFragment.newInstance(restaurant)).commit();
    }

    private int isExistMenu(String menu) {

        for (int i = 0; i < mCartList.size(); i++) {
            if (mCartList.get(i).getMenu().getName().equals(menu))
                return i;
        }

        return -1;
    }

    private int calculateTotalCnt() {
        int sum = 0;
        for (int i = 0; i < mCartList.size(); i++) {
            sum += mCartList.get(i).getCount();
        }

        return sum;
    }
}
