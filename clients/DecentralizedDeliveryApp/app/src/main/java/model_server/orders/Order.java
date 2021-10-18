package model_server.orders;

import com.google.gson.annotations.SerializedName;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Order {
    private Where where;
    private String who;
    private String order_at;
    private String cooking_at;
    private String deliver_at;
    private String complete_at;
    private int state;
    private int id;
    @SerializedName("purchased")
    private List<Purchased> purchasedList;
}
