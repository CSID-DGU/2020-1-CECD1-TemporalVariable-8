package temporary.variable.android.decentralizeddeliveryapp;

import android.content.Context;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.widget.SearchView;
import androidx.appcompat.widget.Toolbar;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;

import java.util.ArrayList;
import java.util.List;

import model.Restaurant;

public class SearchFragment extends Fragment {
    private RecyclerView mRecyclerView;
    private SearchRecyclerAdapter mSearchRecyclerAdapter;
    private MainActivity mContext;

    private List<Restaurant> mRestaurantList;
    private List<Restaurant> mSearchResultList;

    private Toolbar mToolbar;
    private SearchView mSearchView;
    private SwipeRefreshLayout mSwipeRefreshLayout;

    SearchView.OnQueryTextListener mOnQueryTextListener = new SearchView.OnQueryTextListener() {
        @Override
        public boolean onQueryTextSubmit(String query) {

            return true;
        }

        @Override
        public boolean onQueryTextChange(String newText) {
            mSearchResultList = new ArrayList<>();

            for (int i = 0; i < mRestaurantList.size(); i++) {
                if(mRestaurantList.get(i).getCategory().equals(newText)) {
                    mSearchResultList.add(mRestaurantList.get(i));
                }
            }

            mSearchRecyclerAdapter.setRestaurantList(mSearchResultList);
            return true;
        }
    };

    public SearchFragment() {

    }

    @Override
    public void onAttach(Context context) {
        super.onAttach(context);
        mContext = (MainActivity) context;
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = inflater.inflate(R.layout.fragment_search, container, false);


        mRestaurantList = new ArrayList<>();
        List<Integer> imageList1 = new ArrayList<>();
        List<Integer> imageList2 = new ArrayList<>();
        List<Integer> imageList3 = new ArrayList<>();
        List<Integer> imageList4 = new ArrayList<>();
        List<Integer> imageList5 = new ArrayList<>();
        imageList1.add(R.drawable.chicken_2);
        imageList1.add(R.drawable.chicken_3);
        imageList1.add(R.drawable.chicken_4);

        imageList2.add(R.drawable.pizza_1);
        imageList2.add(R.drawable.pizza_2);
        imageList2.add(R.drawable.pizza_3);

        imageList3.add(R.drawable.dbk_1);
        imageList3.add(R.drawable.dbk_2);
        imageList3.add(R.drawable.dbk_3);

        imageList4.add(R.drawable.jokbo_1);
        imageList4.add(R.drawable.jokbo_2);
        imageList4.add(R.drawable.jokbo_3);

        imageList5.add(R.drawable.china_1);
        imageList5.add(R.drawable.china_2);
        imageList5.add(R.drawable.china_3);

        mRestaurantList.add(new Restaurant("임시변수 치킨집","개발자 출신 사장님이 직접 튀기는 치킨!",R.drawable.chicken_logo,imageList1, "치킨","충무로점"));
        mRestaurantList.add(new Restaurant("피자 핫","시카고 피자 전문점입니다.",R.drawable.pizza_logo,imageList2,"피자","명동점"));
        mRestaurantList.add(new Restaurant("우리집 떡볶이","떡볶이 명가 우리집입니다.",R.drawable.dbk_logo,imageList3, "분식","홍대점"));
        mRestaurantList.add(new Restaurant("행복담은족발","행복을 드려요.^^",R.drawable.jokbo_logo,imageList4,"족발", "강남점"));
        mRestaurantList.add(new Restaurant("홍문","정통 중국요리 홍문입니다.",R.drawable.china_logo,imageList5, "중식", "이태원점"));

        mRestaurantList.add(new Restaurant("임시변수 치킨집","개발자 출신 사장님이 직접 튀기는 치킨!",R.drawable.chicken_logo,imageList1, "치킨","충무로점"));
        mRestaurantList.add(new Restaurant("피자 핫","시카고 피자 전문점입니다.",R.drawable.pizza_logo,imageList2,"피자","명동점"));
        mRestaurantList.add(new Restaurant("우리집 떡볶이","떡볶이 명가 우리집입니다.",R.drawable.dbk_logo,imageList3, "분식","홍대점"));
        mRestaurantList.add(new Restaurant("행복담은족발","행복을 드려요.^^",R.drawable.jokbo_logo,imageList4,"족발", "강남점"));
        mRestaurantList.add(new Restaurant("홍문","정통 중국요리 홍문입니다.",R.drawable.china_logo,imageList5, "중식", "이태원점"));

        mRestaurantList.add(new Restaurant("임시변수 치킨집","개발자 출신 사장님이 직접 튀기는 치킨!",R.drawable.chicken_logo,imageList1, "치킨","충무로점"));
        mRestaurantList.add(new Restaurant("피자 핫","시카고 피자 전문점입니다.",R.drawable.pizza_logo,imageList2,"피자","명동점"));
        mRestaurantList.add(new Restaurant("우리집 떡볶이","떡볶이 명가 우리집입니다.",R.drawable.dbk_logo,imageList3, "분식","홍대점"));
        mRestaurantList.add(new Restaurant("행복담은족발","행복을 드려요.^^",R.drawable.jokbo_logo,imageList4,"족발", "강남점"));
        mRestaurantList.add(new Restaurant("홍문","정통 중국요리 홍문입니다.",R.drawable.china_logo,imageList5, "중식", "이태원점"));




        mToolbar = v.findViewById(R.id.search_toolbar);
        mToolbar.inflateMenu(R.menu.search);

        mSearchView = (SearchView) mToolbar.getMenu().findItem(R.id.menu_search).getActionView();
        mSearchView.setOnQueryTextListener(mOnQueryTextListener);

        mRecyclerView = v.findViewById(R.id.search_recyclerView);
        mSearchRecyclerAdapter = new SearchRecyclerAdapter(new ArrayList<Restaurant>());

        mRecyclerView.setLayoutManager(new LinearLayoutManager(getContext()));
        mRecyclerView.setAdapter(mSearchRecyclerAdapter);

        mSwipeRefreshLayout = v.findViewById(R.id.swipe_refresh_layout);
        mSwipeRefreshLayout.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                //새로고침

                mSearchView.setQuery("",false);
                mSwipeRefreshLayout.setRefreshing(false);
            }
        });

        return v;
    }

    static class SearchRecyclerAdapter extends RecyclerView.Adapter<SearchViewHolder> {
        List<Restaurant> restaurantList;

        public SearchRecyclerAdapter(List<Restaurant> list) {
            restaurantList = list;
        }

        @NonNull
        @Override
        public SearchViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
            View v = LayoutInflater.from(parent.getContext()).inflate(R.layout.list_item_restaurant, parent, false);
            return new SearchViewHolder(v);
        }

        @Override
        public void onBindViewHolder(@NonNull SearchViewHolder holder, int i) {
            Restaurant restaurant = restaurantList.get(i);
            holder.name.setText(restaurant.getName());
            holder.description.setText(restaurant.getDescription());

            holder.first.setImageResource(restaurant.getImages().get(0));
            holder.second.setImageResource(restaurant.getImages().get(0));
            holder.third.setImageResource(restaurant.getImages().get(0));
        }

        @Override
        public int getItemCount() {
            return restaurantList.size();
        }

        public void setRestaurantList(List<Restaurant> list) {
            restaurantList = list;
            notifyDataSetChanged();
        }
    }

    static class SearchViewHolder extends RecyclerView.ViewHolder {
        TextView name;
        TextView description;

        ImageView first;
        ImageView second;
        ImageView third;

        public SearchViewHolder(@NonNull View itemView) {
            super(itemView);
            name = itemView.findViewById(R.id.list_item_text_view_name);
            description = itemView.findViewById(R.id.list_item_text_view_description);
            first = itemView.findViewById(R.id.list_item_image_view_first);
            second = itemView.findViewById(R.id.list_item_image_view_second);
            third = itemView.findViewById(R.id.list_item_image_view_third);
        }
    }
}
