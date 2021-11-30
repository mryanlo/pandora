using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.ResourceManager.EventGrid
{
    public partial class Service : ServiceDefinition
    {
        public string Name => "EventGrid";
        public string? ResourceProvider => "Microsoft.EventGrid";
    }
}
