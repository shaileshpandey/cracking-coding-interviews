using System;
using System.IO;
using System.Net;
using System.Threading.Tasks;
using Xunit;
using FluentAssertions;

namespace SetAssociativeCache.Tests
{
    public class FunctionalTests
    {
        [Fact]
        public void TestBasicLruFunctionality ()
        {
            SetAssociativeCache<string,string> nWayCache = new SetAssociativeCache<string, string>(1, 2);
            nWayCache.Set("Eggs", "Ham");
            nWayCache.Set("Sam", "Iam");
            nWayCache.Set("Green", "EggsAndHam");

            nWayCache.Get("Sam").Should().Be("Iam");
            nWayCache.Get("Green").Should().Be("EggsAndHam");

            nWayCache.ContainsKey("Eggs").Should().BeFalse();
            nWayCache.ContainsKey("Sam").Should().BeTrue();
            nWayCache.ContainsKey("Green").Should().BeTrue();
        }
    }
}
