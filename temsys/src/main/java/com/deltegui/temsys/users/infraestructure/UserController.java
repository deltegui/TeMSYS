package com.deltegui.temsys.users.infraestructure;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/dashboard/admin/users")
public class UserController {

    @GetMapping
    public String index() {
        return "user/index";
    }
}
