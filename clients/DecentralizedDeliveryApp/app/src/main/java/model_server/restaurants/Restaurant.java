package model_server.restaurants;

import com.google.gson.annotations.SerializedName;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Restaurant {
    @SerializedName("names")
    private List<Names> names;
}
