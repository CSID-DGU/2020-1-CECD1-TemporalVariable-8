package data_source;

public interface MyListener<T> {
    void onComplete(T result);
}
