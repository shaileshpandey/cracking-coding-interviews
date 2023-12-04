using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading;

/// <summary>
/// A Set-Associative Cache data structure with fixed capacity.
///
/// - Data is structured into `setCount` # of `setSize`-sized sets.
/// - Every possible key is associated with exactly one set via a hashing algorightm
/// - If more items are added to a set than it has capacity for (i.e. > `setSize` items),
///   a replacement victim is chosen from that set using an LRU algorighm.
///
///   NOTE: Part of the exercise is to allow for different kinds of replacement algorithms...
/// </summary>
public class SetAssociativeCache<TKey, TValue>
{
    public int Capacity => this.SetCount * this.SetSize;
    public int SetCount { get; private set; }
    public int SetSize { get; private set; }

    private readonly CacheSet<TKey, TValue>[] sets;

    public SetAssociativeCache(int setCount, int setSize)
    {
        this.SetCount = setCount;
        this.SetSize = setSize;

        // Initialize the sets
        this.sets = new CacheSet<TKey, TValue>[this.SetCount];
        for (int i = 0; i < this.SetCount; i++)
        {
            sets[i] = new CacheSet<TKey, TValue>(setSize);
        }
    }

    /// <summary>
    /// Gets the value associated with `key`. Throws if key not found.
    /// </summary>
    public TValue Get(TKey key)
    {
        int setIndex = this.GetSetIndex(key);
        var set = this.sets[ setIndex ];
        return set.Get(key);
    }

    /// <summary>
    /// Adds the `key` to the cache with the associated value, or overwrites the existing key. 
    /// If adding would exceed capacity, an existing key is chosen to replace using an LRU algorithm
    /// (NOTE: It is part of this exercise to allow for more replacement algos)
    /// </summary>
    public void Set(TKey key, TValue value)
    {
        int setIndex = this.GetSetIndex(key);
        var set = this.sets[setIndex];
        set.Set(key, value);
    }

    /// <summary>
    /// Returns `true` if the given `key` is present in the set; otherwise, `false`.
    /// </summary>
    public bool ContainsKey(TKey key)
    {
        int setIndex = this.GetSetIndex(key);
        var set = this.sets[setIndex];
        return set.ContainsKey(key);
    }
    
    /// <summary>
    /// Returns the count of items in the cache
    /// </summary>
    public int GetCount()
    {
        return this.sets.Sum( s => s.Count );
    }

    /// <summary>
    /// Maps a key to a set
    /// </summary>
    private int GetSetIndex(TKey key)
    {
        return Math.Abs(key.GetHashCode()) % this.SetCount;
    }
}
