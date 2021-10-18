package model_server.menus;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Options {
    private int id;
    private String name;
    private int priceAmount;
    private String currency;
}
