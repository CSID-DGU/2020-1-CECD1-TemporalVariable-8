package data_source;

import java.util.List;

import io.reactivex.rxjava3.core.Flowable;
import model_server.menus.Menu;
import model_server.orders.Order;
import model_server.orders.PostOrder;
import model_server.restaurant_detail.RestaurantDetail;
import model_server.restaurants.Restaurant;
import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Path;

public interface RestaurantApi {
    String BASEURL = "http://ddiggo.iptime.org/";

    @GET("/api/restaurants")
    Flowable<Restaurant> getRestaurantList();

    @GET("/api/restaurants/{name}/menus")
    Flowable<Menu> getMenuList(@Path(value = "name", encoded = true) String name);

    @GET("/api/restaurants/{name}/orders")
    Flowable<List<Order>> getOrderList(@Path(value = "name", encoded = true) String name);

    @GET("/api/restaurants/{name}/")
    Flowable<RestaurantDetail> getRestaurantDetail(@Path(value = "name", encoded = true) String name);

    @POST("/api/restaurants/{name}/order")
    Call<Void> postOrder(@Path(value = "name", encoded = true) String name, @Body PostOrder order);
}
