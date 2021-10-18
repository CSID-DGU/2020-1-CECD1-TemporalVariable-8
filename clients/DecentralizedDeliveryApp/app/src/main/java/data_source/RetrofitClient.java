package data_source;

import android.content.Context;

import hu.akarnokd.rxjava3.retrofit.RxJava3CallAdapterFactory;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class RetrofitClient {
    private static RetrofitClient instance;
    private static Retrofit retrofit;

    private RetrofitClient(){}

    public static RetrofitClient getInstance(Context context) {
        if(instance == null)
            instance = new RetrofitClient();
        return instance;
    }

    public Retrofit getClient() {
        if(retrofit == null) {
            retrofit = new Retrofit.Builder()
                    .baseUrl(RestaurantApi.BASEURL)
                    .addConverterFactory(GsonConverterFactory.create())
                    .addCallAdapterFactory(RxJava3CallAdapterFactory.create())
                    .build();
        }
        return retrofit;
    }
}
