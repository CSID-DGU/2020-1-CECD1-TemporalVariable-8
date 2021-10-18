package model;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Order implements Serializable {
    private Menu menu;
    private Integer count;
    private Integer orderPrice;
}
