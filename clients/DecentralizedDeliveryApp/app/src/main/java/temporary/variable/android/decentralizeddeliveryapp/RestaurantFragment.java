package temporary.variable.android.decentralizeddeliveryapp;

import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

import data_source.Repository;
import data_source.RestaurantApi;
import de.hdodenhof.circleimageview.CircleImageView;
import model.Restaurant;
import model_server.restaurant_detail.RestaurantDetail;
import model_server.restaurants.Names;
import restaurant_detail.RestaurantDetailActivity;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class RestaurantFragment extends Fragment {
    private RecyclerView mRecyclerView;
    private MainActivity mContext;

    private List<Restaurant> mRestaurantList;

    @Override
    public void onAttach(Context context) {
        super.onAttach(context);
        mContext = (MainActivity) context;
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = inflater.inflate(R.layout.fragment_recycler_view, container, false);

//        mRestaurantList = new ArrayList<>();
//        List<Integer> imageList1 = new ArrayList<>();
//        List<Integer> imageList2 = new ArrayList<>();
//        List<Integer> imageList3 = new ArrayList<>();
//        List<Integer> imageList4 = new ArrayList<>();
//        List<Integer> imageList5 = new ArrayList<>();
//        imageList1.add(R.drawable.chicken_2);
//        imageList1.add(R.drawable.chicken_3);
//        imageList1.add(R.drawable.chicken_4);
//
//        imageList2.add(R.drawable.pizza_1);
//        imageList2.add(R.drawable.pizza_2);
//        imageList2.add(R.drawable.pizza_3);
//
//        imageList3.add(R.drawable.dbk_1);
//        imageList3.add(R.drawable.dbk_2);
//        imageList3.add(R.drawable.dbk_3);
//
//        imageList4.add(R.drawable.jokbo_1);
//        imageList4.add(R.drawable.jokbo_2);
//        imageList4.add(R.drawable.jokbo_3);
//
//        imageList5.add(R.drawable.china_1);
//        imageList5.add(R.drawable.china_2);
//        imageList5.add(R.drawable.china_3);

//        mRestaurantList.add(new Restaurant("임시변수 치킨집","개발자 출신 사장님이 직접 튀기는 치킨!",R.drawable.chicken_logo,imageList1, "치킨","충무로점"));
//        mRestaurantList.add(new Restaurant("피자 핫","시카고 피자 전문점입니다.",R.drawable.pizza_logo,imageList2,"피자","명동점"));
//        mRestaurantList.add(new Restaurant("우리집 떡볶이","떡볶이 명가 우리집입니다.",R.drawable.dbk_logo,imageList3, "분식","홍대점"));
//        mRestaurantList.add(new Restaurant("행복담은족발","행복을 드려요.^^",R.drawable.jokbo_logo,imageList4,"족발", "강남점"));
//        mRestaurantList.add(new Restaurant("홍문","정통 중국요리 홍문입니다.",R.drawable.china_logo,imageList5, "중식", "이태원점"));

        mRecyclerView = v.findViewById(R.id.restaurant_recyclerView);
//        RestaurantRecyclerAdapter mRestaurantRecyclerAdapter = new RestaurantRecyclerAdapter(mRestaurantList);
        RestaurantRecyclerAdapter mRestaurantRecyclerAdapter = new RestaurantRecyclerAdapter(Repository.getInstance(getContext()).getRestaurantList());
        mRecyclerView.setLayoutManager(new LinearLayoutManager(getContext()));
        mRecyclerView.setAdapter(mRestaurantRecyclerAdapter);
        return v;
    }

    class RestaurantRecyclerAdapter extends RecyclerView.Adapter<RestaurantViewHolder> {
//        List<Restaurant> restaurantList;
        List<Names> restaurantList;
        String description = "empty";
        //        public RestaurantRecyclerAdapter(List<Restaurant> list) {
//            restaurantList = list;
//        }
        public RestaurantRecyclerAdapter(List<Names> list) {
            Log.d("MYTAG", "size : " + list.size());
            restaurantList = list;
        }

        @NonNull
        @Override
        public RestaurantViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
            View v = LayoutInflater.from(parent.getContext()).inflate(R.layout.list_item_restaurant, parent, false);
            return new RestaurantViewHolder(v);
        }

        @Override
        public void onBindViewHolder(@NonNull RestaurantViewHolder holder, int i) {
//            final Restaurant restaurant = restaurantList.get(i);
            List<String> path = restaurantList.get(i).getPath();
            String name = path.get(0);

            StringBuilder sb = new StringBuilder();
            for (int j = 0; j < path.size(); j++) {
                sb.append(path.get(j)).append(" ");
            }


            Repository.getInstance(getContext()).fetchDescription(name, (description) -> {
                holder.name.setText(sb.toString().trim());
                holder.description.setText(description);
                holder.logo.setImageResource(R.drawable.china_logo);
                holder.first.setImageResource(R.drawable.china_1);
                holder.second.setImageResource(R.drawable.china_2);
                holder.third.setImageResource(R.drawable.china_3);
                holder.itemView.setOnClickListener(view -> {
                    Intent intent = new Intent(mContext, RestaurantDetailActivity.class);
                    intent.putExtra("restaurant", (Serializable) path);
                    startActivity(intent);
                });
            });


//            holder.name.setText(restaurant.getName() + " " + restaurant.getLocation());
//            holder.description.setText(restaurant.getDescription());
//
//            holder.logo.setImageResource(restaurant.getLogoImage());
//
//            holder.first.setImageResource(restaurant.getImages().get(0));
//            holder.second.setImageResource(restaurant.getImages().get(1));
//            holder.third.setImageResource(restaurant.getImages().get(2));
//
//            holder.itemView.setOnClickListener(view -> {
//                Intent intent = new Intent(mContext, RestaurantDetailActivity.class);
//                intent.putExtra("restaurant", restaurant);
//                startActivity(intent);
//            });

        }

        @Override
        public int getItemCount() {
            return restaurantList.size();
        }

        public void setRestaurantList(List<Names> list) {
            restaurantList = list;
            notifyDataSetChanged();
        }
    }

    class RestaurantViewHolder extends RecyclerView.ViewHolder {
        TextView name;
        TextView description;

        CircleImageView logo;

        ImageView first;
        ImageView second;
        ImageView third;

        public RestaurantViewHolder(@NonNull View itemView) {
            super(itemView);
            name = itemView.findViewById(R.id.list_item_text_view_name);
            description = itemView.findViewById(R.id.list_item_text_view_description);

            logo = itemView.findViewById(R.id.list_item_restaurant_logo_image);
            first = itemView.findViewById(R.id.list_item_image_view_first);
            second = itemView.findViewById(R.id.list_item_image_view_second);
            third = itemView.findViewById(R.id.list_item_image_view_third);
        }
    }
}
