package model_server.menus;


import com.google.gson.annotations.SerializedName;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Menu {
    @SerializedName("owner")
    private Owner ower;
    @SerializedName("menus")
    private List<Menus> menuList;
}
