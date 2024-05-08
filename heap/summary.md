# heap


### Problems

#### 295. Get the median of an ordered stream

* Key Point

  * Keep a max heap and min heap 
  * May use the go built-in heap package

      * The `up` and `down` operation is a typical `swim` and `sink` operation of a heap
      
* Principle

  * Using the heap sort. (if the stream is ordered, no need to order again)
