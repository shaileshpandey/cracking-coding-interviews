/// <summary>
/// An internal data structure representing a single item in an N-Way Set-Associative Cache
/// </summary>
internal class CacheItem<TKey, TValue>
{
    public TKey Key;
    public TValue Value;

    public CacheItem(TKey key, TValue value)
    {
        this.Key = key;
        this.Value = value;
    }
}
