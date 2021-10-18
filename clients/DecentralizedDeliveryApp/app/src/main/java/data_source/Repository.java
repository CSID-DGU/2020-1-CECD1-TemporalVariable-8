package data_source;

import android.content.Context;
import android.util.Log;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

import hu.akarnokd.rxjava3.retrofit.RxJava3CallAdapterFactory;
import io.reactivex.rxjava3.android.schedulers.AndroidSchedulers;
import io.reactivex.rxjava3.disposables.CompositeDisposable;
import io.reactivex.rxjava3.schedulers.Schedulers;
import model.OrderState;
import model.Review;
import model_server.menus.Menu;
import model_server.menus.Menus;
import model_server.orders.Order;
import model_server.orders.PostOrder;
import model_server.restaurant_detail.RestaurantDetail;
import model_server.restaurants.Names;
import model_server.restaurants.Restaurant;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import temporary.variable.android.decentralizeddeliveryapp.RestaurantFragment;

public class Repository {
    private static Repository sRepository;
    private static List<Names> sRestaurantList;
    private static List<Order> sOrderList;
    private static List<Menus> sMenuList;
    private static List<Review> sReviewList;

    public static CompositeDisposable mCompositeDisposable;

    private Repository() {
        mCompositeDisposable = new CompositeDisposable();
    }

    public static Repository getInstance(Context context) {
        if (sRepository == null)
            sRepository = new Repository();

        return sRepository;
    }
    public void sendOrder(String name, PostOrder postOrder){
        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl(RestaurantApi.BASEURL)
                .addConverterFactory(GsonConverterFactory.create())
                .build();
        final Repository thiso = this;
        RestaurantApi restaurantApi = retrofit.create(RestaurantApi.class);
        restaurantApi.postOrder(name, postOrder).enqueue(new Callback<Void>() {
            @Override
            public void onResponse(Call<Void> call, Response<Void> response) {
                Log.d("TAG", "동국반점 주문: ");
            }

            @Override
            public void onFailure(Call<Void> call, Throwable t) {

                Log.d("MYTAG", "sendOrder fail");
            }
        });
    }
    public void fetchRestaurantList(MyListener<Boolean> callback) {

        if(sRestaurantList == null) {
            sRestaurantList = new ArrayList<>();
            Retrofit retrofit = new Retrofit.Builder()
                    .baseUrl(RestaurantApi.BASEURL)
                    .addConverterFactory(GsonConverterFactory.create())
                    .addCallAdapterFactory(RxJava3CallAdapterFactory.create())
                    .build();

            RestaurantApi restaurantApi = retrofit.create(RestaurantApi.class);
            mCompositeDisposable.add(restaurantApi.getRestaurantList()
                    .subscribeOn(Schedulers.computation())
                    .observeOn(AndroidSchedulers.mainThread())
                    .subscribe(data -> {
                        Log.d("MYTAG", "getRestaurantList : " + data.getNames().size());
                        sRestaurantList = data.getNames();
                        callback.onComplete(null);
                    }));
        }


//        try {
//            Response<Restaurant> response = restaurantApi.getRestaurantList().execute();
//            sRestaurantList = response.body().getNames();
//        } catch (IOException e) {
//            e.printStackTrace();
//        }

//        restaurantApi.getRestaurantList().enqueue(new Callback<model_server.restaurants.Restaurant>() {
//            @Override
//            public void onResponse(Call<model_server.restaurants.Restaurant> call, Response<model_server.restaurants.Restaurant> response) {
//                sRestaurantList = response.body().getNames();
////                    Log.d("MYTAG", response.body().getNames().get(0).getPath().get(0));
////                    Log.d("MYTAG", response.body().getNames().get(1).getPath().get(1));
//            }
//
//            @Override
//            public void onFailure(Call<model_server.restaurants.Restaurant> call, Throwable t) {
//
//            }
//        });


    }

    public void fetchMenuList(String name, MyListener<Boolean> callback) {

        if(sMenuList == null) {
            sMenuList = new ArrayList<>();
            Retrofit retrofit = new Retrofit.Builder()
                    .baseUrl(RestaurantApi.BASEURL)
                    .addConverterFactory(GsonConverterFactory.create())
                    .addCallAdapterFactory(RxJava3CallAdapterFactory.create())
                    .build();

            RestaurantApi restaurantApi = retrofit.create(RestaurantApi.class);
            mCompositeDisposable.add(restaurantApi.getMenuList(name)
                    .subscribeOn(Schedulers.computation())
                    .observeOn(AndroidSchedulers.mainThread())
                    .subscribe(data -> {
                        Log.d("MYTAG", "getMenuList : " + data);
                        sMenuList = data.getMenuList();
                        callback.onComplete(null);
                    }));
        }


//        restaurantApi.getMenuList(name).enqueue(new Callback<model_server.menus.Menu>() {
//            @Override
//            public void onResponse(Call<model_server.menus.Menu> call, Response<model_server.menus.Menu> response) {
//                sMenuList = response.body().getMenuList();
//
////                    Log.d("MYTAG", "name : " + response.body().getMenuList().get(0).getName());
////                    Log.d("MYTAG", response.body().getMenuList().get(0).getDescription());
////                    Log.d("MYTAG", "price : " + response.body().getMenuList().get(0).getPriceAmount());
////                    Log.d("MYTAG", "options : " + response.body().getMenuList().get(0).getOptions().get(0).getName());
////                    Log.d("MYTAG", "options priceAmount : " + response.body().getMenuList().get(0).getOptions().get(0).getPriceAmount());
//            }
//
//            @Override
//            public void onFailure(Call<model_server.menus.Menu> call, Throwable t) {
//
//            }
//        });


    }

    public void fetchOrderList(String name, MyListener<Boolean> callback) {

        sOrderList = new ArrayList<>();
        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl(RestaurantApi.BASEURL)
                .addConverterFactory(GsonConverterFactory.create())
                .addCallAdapterFactory(RxJava3CallAdapterFactory.create())
                .build();

        RestaurantApi restaurantApi = retrofit.create(RestaurantApi.class);
        mCompositeDisposable.add(restaurantApi.getOrderList(name)
                .subscribeOn(Schedulers.computation())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribe(data -> {
                    Log.d("MYTAG", "getOrderList : " + data);
                    sOrderList = data;
                    callback.onComplete(null);
                }));


//        restaurantApi.getOrderList(name).enqueue(new Callback<List<Order>>() {
//            @Override
//            public void onResponse(Call<List<Order>> call, Response<List<Order>> response) {
//                sOrderList = response.body();
////                Log.d("MYTAG", "where path : " + response.body().get(0).getWhere().getPath().get(0));
////                Log.d("MYTAG", "state : " + response.body().get(0).getState());
////                Log.d("MYTAG", "order_at : " + response.body().get(0).getOrder_at());
////
////                Log.d("MYTAG", "purchased name : " + response.body().get(0).getPurchasedList().get(0).getName());
////                Log.d("MYTAG", "purchased amount : " + response.body().get(0).getPurchasedList().get(0).getAmount());
//            }
//
//            @Override
//            public void onFailure(Call<List<Order>> call, Throwable t) {
//
//            }
//        });
    }

    public void fetchDescription(String name, MyListener<String> callback) {
        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl(RestaurantApi.BASEURL)
                .addConverterFactory(GsonConverterFactory.create())
                .addCallAdapterFactory(RxJava3CallAdapterFactory.create())
                .build();

        RestaurantApi restaurantApi = retrofit.create(RestaurantApi.class);
        mCompositeDisposable.add(restaurantApi.getRestaurantDetail(name)
                .subscribeOn(Schedulers.computation())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribe(data -> {
                    Log.d("MYTAG", "getDescription : " + data.getDescription());
                    callback.onComplete(data.getDescription());
                }));
    }

    public List<Names> getRestaurantList() {
        return sRestaurantList;
    }

    public List<Menus> getMenuList(String restaurantName) {
        return sMenuList;
    }

    public List<Order> getOrderList() {
        return sOrderList;
    }


}
