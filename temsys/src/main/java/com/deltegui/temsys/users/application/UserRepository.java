package com.deltegui.temsys.users.application;

import com.deltegui.temsys.users.domain.User;

import java.util.List;

public interface UserRepository {
    List<User> getAll();
    void create(User user);
    void update(User user);
    void deleteByName(String name);
}
