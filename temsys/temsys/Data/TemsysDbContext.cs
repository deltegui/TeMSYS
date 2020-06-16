using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using System.Configuration;

namespace temsys {
    public class TemsysDbContext : DbContext {
        public DbSet<IdentityUser> Users { get; set; }

        public TemsysDbContext(DbContextOptions<TemsysDbContext> options) : base(options) { }

        protected override void OnModelCreating(ModelBuilder modelBuilder) {
            base.OnModelCreating(modelBuilder);
        }
    }
}
