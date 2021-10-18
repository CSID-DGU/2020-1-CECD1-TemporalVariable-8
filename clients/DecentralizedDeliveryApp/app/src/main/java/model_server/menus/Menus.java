package model_server.menus;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Menus {
    private int id;
    private String name;
    private String description;
    private int priceAmount;
    private String currency;
    private List<Options> options;
}
