package model;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class Menu implements Serializable {
    private int id;
    private String name;
    private Integer price;
    private String description;
    private Integer image;
//    private String category;
//    private Option option;
}
