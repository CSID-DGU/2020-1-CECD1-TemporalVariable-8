package model_server.orders;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Purchased {
    private int menu;
    private int count;
    private String name;
    private int amount;
}
