using System;

namespace Domain.Core {
    public class Maybe<T> {
        private T value;
        private bool filled;

        private Maybe(T value) {
            this.value = value;
            this.filled = true;
        }

        private Maybe() {
            this.value = default(T);
            this.filled = false;
        }

        public static Maybe<T> Fill(T value) => new Maybe<T>(value);
        public static Maybe<T> Empty() => new Maybe<T>();

        public T Unwrap() {
            if (! this.filled) {
                throw new NullReferenceException();
            }
            return this.value;
        }

        public bool IsFilled() {
            return this.filled;
        }

        public Maybe<T> IfPresent(Action<T> consumer) {
            if (this.IsFilled()) {
                consumer(this.value);
            }
            return this;
        }

        public Maybe<T> IfEmpty(Action consumer) {
            if (! this.IsFilled()) {
                consumer();
            }
            return this;
        }
    }

    public abstract class Result<TOk, TErr> {
        public class Ok : Result<TOk, TErr> {
            public TOk Value { get; }
            public Ok(TOk value) => this.Value = value;
        }

        public class Err : Result<TOk, TErr> {
            public TErr Value { get; }
            public Err(TErr value) => this.Value = value;
        }

        public static Ok WithValue(TOk value) => new Ok(value);
        public static Err WithError(TErr value) => new Err(value);
    }
}