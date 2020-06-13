using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace temsys.Services {
    public class RepositoryException : Exception {
        public RepositoryException(string message) : base(message) { }
    }
}
