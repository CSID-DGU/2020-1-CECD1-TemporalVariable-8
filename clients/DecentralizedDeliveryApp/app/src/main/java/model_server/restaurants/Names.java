package model_server.restaurants;

import com.google.gson.annotations.SerializedName;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Names {
    @SerializedName("path")
    private List<String> path;

    @SerializedName("path1")
    private String path1;
//    @SerializedName("path2")
//    private String path2;

}
