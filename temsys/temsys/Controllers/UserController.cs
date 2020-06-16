using System;
using System.Diagnostics;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Linq;
using Microsoft.AspNetCore.Mvc;
using temsys.Models;
using Microsoft.AspNetCore.Identity;
using temsys.Services;
using Microsoft.EntityFrameworkCore;

namespace temsys.Controllers {
    public class UserController : Controller {
        private readonly TemsysDbContext repository;

        public UserController(TemsysDbContext repository) {
            this.repository = repository;
        }

        public IActionResult Index() {
            return View();
        }

        [HttpPost]
        [ValidateAntiForgeryToken]
        public IActionResult Login(LoginFormViewModel model) {
            if(!ModelState.IsValid) {
                return View("Index", model);
            }
            var users = from u in repository.Users where u.UserName == model.Username select u;
            if(!users.Any()) {
                ModelState.AddModelError(string.Empty, "User does not exist");
                return View("Index");
            }
            var user = users.First();
            if(user.PasswordHash != model.Password) {
                ModelState.AddModelError(string.Empty, "Incorrect password");
                return View("Index");
            }
            return Content($"Name: {user.UserName}, Password: {user.PasswordHash}");
        }
    }
}