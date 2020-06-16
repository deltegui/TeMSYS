using System;

namespace temsys.Services {
    public class RepositoryException : Exception {
        public RepositoryException(string message) : base(message) { }
    }
}
