package restaurant_detail;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.DividerItemDecoration;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.google.android.material.bottomsheet.BottomSheetBehavior;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

import data_source.Repository;
import model.Menu;
import model.Order;
import model.Restaurant;
import model_server.menus.Menus;
import temporary.variable.android.decentralizeddeliveryapp.R;
import util.Util;

public class RestaurantMenuFragment extends Fragment {
    private RecyclerView mRecyclerView;
    private MenuRecyclerAdapter mRestaurantRecyclerAdapter;

    private List<Menu> mMenuList;
    //    private Restaurant mRestaurant;
    private String mRestaurantName;

    RestaurantMenuFragment() {
    }

//    public static Fragment newInstance(Restaurant restaurant) {
//
//        Bundle args = new Bundle();
//        Fragment fragment = new RestaurantMenuFragment();
//
//        args.putSerializable("restaurant", restaurant);
//        fragment.setArguments(args);
//
//        return fragment;
//    }

    public static Fragment newInstance(String restaurantName) {

        Bundle args = new Bundle();
        Fragment fragment = new RestaurantMenuFragment();

        args.putSerializable("restaurant", (Serializable) restaurantName);
        fragment.setArguments(args);

        return fragment;
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = inflater.inflate(R.layout.fragment_restaurant_menu, container, false);

        mRecyclerView = v.findViewById(R.id.recyclerView);

        if (getArguments() != null) {
//            mRestaurant = (Restaurant) getArguments().getSerializable("restaurant");
            mRestaurantName = (String) getArguments().getSerializable("restaurant");
        }

        mMenuList = new ArrayList<>();

        List<Menus> list = Repository.getInstance(getContext()).getMenuList(mRestaurantName);
        if (mRestaurantName.equals("동국반점")) {
            for (int i = 0; i < list.size(); i++) {
                mMenuList.add(new Menu(list.get(i).getId(), list.get(i).getName(),list.get(i).getPriceAmount(),list.get(i).getDescription(),R.drawable.china_1));
            }

        } else if (mRestaurantName.equals("동국반점 2호점")) {
            for (int i = 0; i < list.size(); i++) {
                mMenuList.add(new Menu(list.get(i).getId(),list.get(i).getName(),list.get(i).getPriceAmount(),list.get(i).getDescription(),R.drawable.china_2));
            }
        }

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

        mRestaurantRecyclerAdapter = new MenuRecyclerAdapter(mMenuList);

        mRecyclerView.setLayoutManager(new LinearLayoutManager(getContext()));
        mRecyclerView.addItemDecoration(new DividerItemDecoration(getContext(), LinearLayoutManager.VERTICAL));

        mRecyclerView.setAdapter(mRestaurantRecyclerAdapter);
        return v;
    }

    class MenuRecyclerAdapter extends RecyclerView.Adapter<MenuViewHolder> {
        List<Menu> menuList;

        public MenuRecyclerAdapter(List<Menu> list) {
            menuList = list;
        }

        @NonNull
        @Override
        public MenuViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
            View v = LayoutInflater.from(parent.getContext()).inflate(R.layout.list_item_menu, parent, false);
            return new MenuViewHolder(v);
        }

        @Override
        public void onBindViewHolder(@NonNull MenuViewHolder holder, int i) {
            final Menu menu = menuList.get(i);
            holder.name.setText(menu.getName());
            holder.price.setText(Util.numberToCommaFormat(menu.getPrice()) + "원");
            holder.description.setText(menu.getDescription());
            // TODO : image
            holder.menuImage.setImageResource(menu.getImage());


            holder.itemView.setOnClickListener(view -> {
                RestaurantDetailActivity.mMenu = menu;

                RestaurantDetailActivity.cnt = 1;
                RestaurantDetailActivity.price = menu.getPrice();

                int state = RestaurantDetailActivity.mBottomSheetBehavior.getState();
//                Log.d("MYTAG", "State : " + state);

                RestaurantDetailActivity.mMenuName.setText(menu.getName());
                RestaurantDetailActivity.mPrice.setText(Util.numberToCommaFormat(RestaurantDetailActivity.price) + "원");
                RestaurantDetailActivity.mCount.setText(String.valueOf(RestaurantDetailActivity.cnt));

                if (state == BottomSheetBehavior.STATE_HIDDEN) {
                    RestaurantDetailActivity.mBottomSheetBehavior.setState(BottomSheetBehavior.STATE_EXPANDED);

                } else if (state == BottomSheetBehavior.STATE_COLLAPSED) {
                    RestaurantDetailActivity.mBottomSheetBehavior.setPeekHeight(RestaurantDetailActivity.mLinearLayout.getHeight());
//                        Log.d("MYTAG", "Height : " + mLinearLayout.getHeight());

                } else if (state == BottomSheetBehavior.STATE_EXPANDED) {

                }
            });
        }

        @Override
        public int getItemCount() {
            return menuList.size();
        }

        public void setMenuList(List<Menu> list) {
            menuList = list;
            notifyDataSetChanged();
        }
    }

    class MenuViewHolder extends RecyclerView.ViewHolder {
        TextView name;
        TextView price;
        TextView description;

        ImageView menuImage;


        public MenuViewHolder(@NonNull View itemView) {
            super(itemView);
            name = itemView.findViewById(R.id.menu_name);
            price = itemView.findViewById(R.id.menu_price);
            description = itemView.findViewById(R.id.menu_description);
            menuImage = itemView.findViewById(R.id.menu_image);
        }
    }
}
