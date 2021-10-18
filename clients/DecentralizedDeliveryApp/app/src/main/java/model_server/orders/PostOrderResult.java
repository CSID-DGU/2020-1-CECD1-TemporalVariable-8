package model_server.orders;

import com.google.gson.annotations.SerializedName;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class PostOrderResult {
    @SerializedName("order_id")
    private int orderId;
}
