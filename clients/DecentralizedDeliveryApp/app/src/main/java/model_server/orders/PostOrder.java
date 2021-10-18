package model_server.orders;

import com.google.gson.annotations.SerializedName;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class PostOrder {
    @SerializedName("fields")
    private List<OrderField> fields;
}