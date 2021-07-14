using System;
using Xunit;
using Domain.Core;

namespace domain_test {
    public class ResultTest {
        [Fact]
        public void ResultCanBeFilledWithValue() {
            var result = Result<string, string>.WithValue("Hola");
            Assert.True(result.Value == "Hola");
        }

        [Fact]
        public void ResultCanBeFilledWithError() {
            var result = Result<string, string>.WithError("Error");
            Assert.True(result.Value == "Error");
        }

        [Fact]
        public void ResultCanBeEitherValueOrError() {
            Result<string, string> result = Result<string, string>.WithValue("Hola");
            var value = result switch {
                Result<string, string>.Ok _ => "ok",
                Result<string, string>.Err _ => "error",
                _ => "default",
            };
            Assert.True(value == "ok");
        }
    }

    public class MaybeTest {
        [Fact]
        public void ResultCanBeFilled() {
            var maybe = Maybe<string>.Fill("hola");
            maybe
                .IfPresent((str) => Assert.True(str == "hola"))
                .IfEmpty(() => Assert.True(false));
        }

        [Fact]
        public void ResultCanBeEmpty() {
            var maybe = Maybe<string>.Empty();
            maybe
                .IfPresent((str) => Assert.True(false))
                .IfEmpty(() => Assert.True(true));
        }

        [Fact]
        public void ResultCanBeUnwrapped() {
            var maybe = Maybe<int>.Empty();
            Assert.ThrowsAny<Exception>(() => maybe.Unwrap());
        }
    }
}
