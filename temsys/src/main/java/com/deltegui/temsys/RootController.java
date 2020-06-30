package com.deltegui.temsys;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.servlet.view.RedirectView;

@Controller
public class RootController {

    @GetMapping("/")
    public RedirectView redirectToDashboard() {
        return new RedirectView("/dashboard/status");
    }
}
