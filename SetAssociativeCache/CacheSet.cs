using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading;

/// <summary>
/// An internal data structure representing one set in a N-Way Set-Associative Cache
/// </summary>
internal class CacheSet<TKey, TValue>
{
    public int Count { get; private set; }

    private readonly int capacity;
    private readonly CacheItem<TKey, TValue>[] store;
    private readonly LinkedList<TKey> usageTracker;

    public CacheSet(int capacity)
    {
        this.capacity = capacity;
        this.usageTracker = new LinkedList<TKey>();
        this.store = new CacheItem<TKey, TValue>[capacity];
    }    

    /// <summary>
    /// Gets the value associated with `key`. Throws if key not found.
    /// </summary>
    public TValue Get(TKey key)
    {
        // If the key is present, update the usage tracker
        if (this.ContainsKey(key))
        {
            this.RecordUsage(key);
        }
        else
        {
            throw new KeyNotFoundException($"The key '{key}' was not found");
        }

        return this.store[this.FindIndexOfKey(key)].Value;
    }

    /// <summary>
    /// Adds the `key` to the cache with the associated value, or overwrites the existing key. 
    /// If adding would exceed capacity, an existing key is chosen to replace using an LRU algorithm
    /// (NOTE: It is part of this exercise to allow for more replacement algos)
    /// </summary>
    public void Set(TKey key, TValue value)
    {
        if ( this.ContainsKey(key) )
        {            
            this.store[this.FindIndexOfKey(key)].Value = value;
        }
        else
        {
            int indexToSet;
            // If the set is at it's capacty
            if (this.Count == this.capacity)
            {
                // Choose the Least-Recently-Used (LRU) item to replace, which will be at the tail of the usage tracker
                // TODO: Factor this logic out to allow for custom replacement algos
                var keyToReplace = this.usageTracker.Last.Value;
                indexToSet = this.FindIndexOfKey(keyToReplace);

                // Remove the existing key
                this.RemoveKey(keyToReplace);
            }
            else
            {
                indexToSet = this.Count;
            }

            this.store[indexToSet] = new CacheItem<TKey, TValue>(key, value);
            this.Count++;

        }

        this.RecordUsage(key);
    }

    /// <summary>
    /// Returns `true` if the given `key` is present in the set; otherwise, `false`.
    /// </summary>
    public bool ContainsKey(TKey key)
    {
        return this.FindIndexOfKey(key) >= 0;
    }

    internal void RemoveKey(TKey key)
    {
        var indexOfKey = this.FindIndexOfKey(key);
        if (indexOfKey >= 0)
        {
            this.usageTracker.Remove(key);
            this.store[indexOfKey] = null;
            this.Count--;
        }
    }

    private int FindIndexOfKey(TKey key)
    {
        for (int i = 0; i < this.Count; i++)
        {
            if (this.store[i] != null && this.store[i].Key.Equals(key)) return i;
        }
        return -1;
    }

    private void RecordUsage(TKey key)
    {
        this.usageTracker.Remove(key);
        this.usageTracker.AddFirst(key);
    }
}
