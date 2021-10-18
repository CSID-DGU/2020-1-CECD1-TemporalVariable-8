//package restaurant_detail;
//
//import android.app.Activity;
//import android.content.Context;
//import android.content.Intent;
//import android.graphics.Color;
//import android.os.Bundle;
//import android.util.Log;
//import android.view.LayoutInflater;
//import android.view.View;
//import android.view.ViewGroup;
//import android.widget.Button;
//import android.widget.ImageView;
//import android.widget.LinearLayout;
//import android.widget.TextView;
//import android.widget.Toast;
//
//import androidx.annotation.NonNull;
//import androidx.annotation.Nullable;
//import androidx.appcompat.widget.Toolbar;
//import androidx.fragment.app.Fragment;
//import androidx.recyclerview.widget.DividerItemDecoration;
//import androidx.recyclerview.widget.LinearLayoutManager;
//import androidx.recyclerview.widget.RecyclerView;
//
//import com.google.android.material.appbar.CollapsingToolbarLayout;
//import com.google.android.material.bottomsheet.BottomSheetBehavior;
//import com.google.android.material.floatingactionbutton.FloatingActionButton;
//
//import java.io.Serializable;
//import java.util.ArrayList;
//import java.util.List;
//
//import cart.CartActivity;
//import model.Menu;
//import model.Order;
//import model.Restaurant;
//import temporary.variable.android.decentralizeddeliveryapp.R;
//import util.Util;
//
//public class RestaurantDetailFragment extends Fragment {
//    private RecyclerView mRecyclerView;
//    private MenuRecyclerAdapter mRestaurantRecyclerAdapter;
//    private RestaurantDetailActivity mContext;
//
//    private List<Menu> mMenuList;
//    private Restaurant mRestaurant;
//
//    private CollapsingToolbarLayout mCollapsingToolbarLayout;
//    private Toolbar mToolbar;
//    private FloatingActionButton mCartFab;
//
//    private BottomSheetBehavior mBottomSheetBehavior;
//
//    private LinearLayout mLinearLayout;
//
//    private ImageView mMinus;
//    private ImageView mPlus;
//
//    private TextView mPrice;
//    private TextView mCount;
//    private TextView mMenuName;
//    private TextView mCartCount;
//    private TextView mSubtitle;
//
//    private Button mButtonAddCart;
//
//    private int cnt = 1;
//    private int price = 0;
//
//    public static List<Order> mCartList;
//    private Menu mMenu;
//
//    public static RestaurantDetailFragment newInstance(Restaurant restaurant) {
//
//        Bundle args = new Bundle();
//        RestaurantDetailFragment fragment = new RestaurantDetailFragment();
//
//        args.putSerializable("restaurant", restaurant);
//        fragment.setArguments(args);
//
//        return fragment;
//    }
//
//    public RestaurantDetailFragment() {
//    }
//
//    @Override
//    public void onAttach(Context context) {
//        super.onAttach(context);
//        mContext = (RestaurantDetailActivity) context;
////        Log.d("MYTAG", "onAttach");
//    }
//
//    @Override
//    public void onResume() {
//        super.onResume();
////        Log.d("MYTAG", "onResume");
//
//        if(mCartList.size() > 0)
//            mCartCount.setText(String.valueOf(mCartList.size()));
//    }
//
//
//    @Nullable
//    @Override
//    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
//        View v = inflater.inflate(R.layout.fragment_restaurant_detail, container, false);
////        Log.d("MYTAG", "onCreateView");
//        mCartList = new ArrayList<>();
//
//        mCartCount = v.findViewById(R.id.textView_cart_count);
//        mMenuName = v.findViewById(R.id.textView_menu_name);
//        mPrice = v.findViewById(R.id.textView_price);
//        mCount = v.findViewById(R.id.textView_cnt);
//        mMinus = v.findViewById(R.id.imageView_cnt_minus);
//        mPlus = v.findViewById(R.id.imageView_cnt_plus);
//        mButtonAddCart = v.findViewById(R.id.btn_add_cart);
//
//        mMinus.setOnClickListener(view -> {
//            if (cnt >= 2) {
//                cnt--;
//                price -= mMenu.getPrice();
//            }
//
//            mPrice.setText(Util.numberToCommaFormat(price));
//            mCount.setText(String.valueOf(cnt));
//        });
//
//        mPlus.setOnClickListener(view -> {
//            cnt++;
//            price += mMenu.getPrice();
//
//            mPrice.setText(Util.numberToCommaFormat(price));
//            mCount.setText(String.valueOf(cnt));
//        });
//
//        mButtonAddCart.setOnClickListener(view -> {
//            Menu menu = new Menu(mMenu.getName(), mMenu.getPrice(), mMenu.getDescription(), mMenu.getImage());
//            mCartList.add(new Order(menu, cnt, price));
//
//            mBottomSheetBehavior.setState(BottomSheetBehavior.STATE_HIDDEN);
//            mCartCount.setText(String.valueOf(mCartList.size()));
//            Toast.makeText(getContext(),"장바구니에 추가되었습니다.",Toast.LENGTH_SHORT).show();
//        });
//
//        mLinearLayout = v.findViewById(R.id.bottom_sheet_layout);
//        mBottomSheetBehavior = BottomSheetBehavior.from(mLinearLayout);
//
//        if (getArguments() != null) {
//            mRestaurant = (Restaurant) getArguments().getSerializable("restaurant");
//        }
//
//        mCartFab = v.findViewById(R.id.fab_cart);
//        mCartFab.setOnClickListener(view -> {
//            Intent intent = new Intent(getContext(), CartActivity.class);
////            intent.putExtra("cartList", (Serializable) mCartList);
//            intent.putExtra("restaurant", (Serializable) mRestaurant);
//            startActivity(intent);
//        });
//
////        Log.d("RestaurantDetailFragment", mRestaurant.getName() + "," + mRestaurant.getCategory());
//
//        mCollapsingToolbarLayout = v.findViewById(R.id.toolbar_layout);
//        mCollapsingToolbarLayout.setTitle(mRestaurant.getName());
//        mCollapsingToolbarLayout.setBackgroundResource(mRestaurant.getImages().get(0));
//        mCollapsingToolbarLayout.setExpandedTitleColor(Color.BLACK);
//
//        mSubtitle = v.findViewById(R.id.textView_toolbar_subtitle);
//        mSubtitle.setText(mRestaurant.getLocation());
//
//        mToolbar = v.findViewById(R.id.toolbar);
//        mToolbar.setTitle(mRestaurant.getName() + " " + mRestaurant.getLocation());
//        mRecyclerView = v.findViewById(R.id.recyclerView);
//
//        mMenuList = new ArrayList<>();
//
//        if (mRestaurant.getName().equals("임시변수 치킨집")) {
//            mMenuList.add(new Menu("객체지향치킨", 14000, "후라이드치킨", R.drawable.chicken_1));
//            mMenuList.add(new Menu("DES치킨", 15000, "핫후라이드치킨", R.drawable.chicken_2));
//            mMenuList.add(new Menu("엔지니어링 간장치킨", 15000, "간장치킨", R.drawable.chicken_3));
//            mMenuList.add(new Menu("RSA치킨", 15000, "양념치킨", R.drawable.chicken_4));
//            mMenuList.add(new Menu("암호핫 네트워크 양념", 16000, "핫양념치킨", R.drawable.small_chicken_3));
//            mMenuList.add(new Menu("심프닭", 16000, "파닭", R.drawable.small_chicken_2));
//            mMenuList.add(new Menu("DES반 RSA반", 16000, "핫후라이드 반 양념 반 ", R.drawable.small_chicken_1));
//        } else if (mRestaurant.getName().equals("피자 핫")) {
//            mMenuList.add(new Menu("클래식 토마토 시카고피자", 15000, "음식설명1", R.drawable.pizza_1));
//            mMenuList.add(new Menu("콰트로 포르마쥬 시카고피자", 16000, "음식설명2", R.drawable.pizza_2));
//            mMenuList.add(new Menu("바질 페스토 시카고피자", 15000, "음식설명3", R.drawable.pizza_3));
//        } else if (mRestaurant.getName().equals("우리집 떡볶이")) {
//            mMenuList.add(new Menu("떡볶이 1인분", 3000, "음식설명1", R.drawable.dbk_1));
//            mMenuList.add(new Menu("치즈떡볶이 1인분", 4000, "음식설명2", R.drawable.dbk_2));
//            mMenuList.add(new Menu("짜장떡볶이 1인분", 4000, "음식설명3", R.drawable.dbk_3));
//        } else if (mRestaurant.getName().equals("행복담은족발")) {
//            mMenuList.add(new Menu("족발보쌈세트", 35000, "음식설명1", R.drawable.jokbo_1));
//            mMenuList.add(new Menu("족발 대", 32000, "음식설명2", R.drawable.jokbo_2));
//            mMenuList.add(new Menu("보쌈 대", 32000, "음식설명3", R.drawable.jokbo_3));
//        } else if (mRestaurant.getName().equals("홍문")) {
//            mMenuList.add(new Menu("짜장면", 5000, "음식설명1", R.drawable.china_1));
//            mMenuList.add(new Menu("탕수육 소", 12000, "음식설명2", R.drawable.china_2));
//            mMenuList.add(new Menu("볶음밥", 6000, "음식설명3", R.drawable.china_3));
//        }
//
//        mRestaurantRecyclerAdapter = new MenuRecyclerAdapter(mMenuList);
//
//        mRecyclerView.setLayoutManager(new LinearLayoutManager(getContext()));
//        mRecyclerView.addItemDecoration(new DividerItemDecoration(getContext(), LinearLayoutManager.VERTICAL));
//        mRecyclerView.setAdapter(mRestaurantRecyclerAdapter);
//        return v;
//    }
//
//    class MenuRecyclerAdapter extends RecyclerView.Adapter<MenuViewHolder> {
//        List<Menu> menuList;
//
//        public MenuRecyclerAdapter(List<Menu> list) {
//            menuList = list;
//        }
//
//        @NonNull
//        @Override
//        public MenuViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
//            View v = LayoutInflater.from(parent.getContext()).inflate(R.layout.list_item_menu, parent, false);
//            return new MenuViewHolder(v);
//        }
//
//        @Override
//        public void onBindViewHolder(@NonNull MenuViewHolder holder, int i) {
//            final Menu menu = menuList.get(i);
//            holder.name.setText(menu.getName());
//            holder.price.setText(Util.numberToCommaFormat(menu.getPrice()) + "원");
//            holder.description.setText(menu.getDescription());
//            holder.menuImage.setImageResource(menu.getImage());
//
//
//            holder.itemView.setOnClickListener(view -> {
//                mMenu = menu;
//
//                cnt = 1;
//                price = menu.getPrice();
//
//                int state = mBottomSheetBehavior.getState();
////                Log.d("MYTAG", "State : " + state);
//
//                mMenuName.setText(menu.getName());
//                mPrice.setText(Util.numberToCommaFormat(price) + "원");
//                mCount.setText(String.valueOf(cnt));
//
//                if (state == BottomSheetBehavior.STATE_HIDDEN) {
//                    mBottomSheetBehavior.setState(BottomSheetBehavior.STATE_EXPANDED);
//
//                } else if (state == BottomSheetBehavior.STATE_COLLAPSED) {
//                    mBottomSheetBehavior.setPeekHeight(mLinearLayout.getHeight());
////                        Log.d("MYTAG", "Height : " + mLinearLayout.getHeight());
//
//                } else if (state == BottomSheetBehavior.STATE_EXPANDED) {
//
//                }
//            });
//        }
//
//        @Override
//        public int getItemCount() {
//            return menuList.size();
//        }
//
//        public void setMenuList(List<Menu> list) {
//            menuList = list;
//            notifyDataSetChanged();
//        }
//    }
//
//    class MenuViewHolder extends RecyclerView.ViewHolder {
//        TextView name;
//        TextView price;
//        TextView description;
//
//        ImageView menuImage;
//
//
//        public MenuViewHolder(@NonNull View itemView) {
//            super(itemView);
//            name = itemView.findViewById(R.id.menu_name);
//            price = itemView.findViewById(R.id.menu_price);
//            description = itemView.findViewById(R.id.menu_description);
//            menuImage = itemView.findViewById(R.id.menu_image);
//        }
//    }
//}
