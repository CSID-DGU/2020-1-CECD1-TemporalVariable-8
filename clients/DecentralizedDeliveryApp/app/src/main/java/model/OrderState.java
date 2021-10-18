package model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class OrderState {
//    private int id;
    private String restaurantName;
    private int logo;
    private Menu menu;
    private String time;
    private String state;
}
