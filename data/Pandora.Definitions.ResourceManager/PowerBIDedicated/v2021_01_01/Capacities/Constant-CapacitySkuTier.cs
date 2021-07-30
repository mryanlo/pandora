using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.PowerBIDedicated.v2021_01_01.Capacities
{
    [ConstantType(ConstantTypeAttribute.ConstantType.String)]
    internal enum CapacitySkuTier
    {
        [Description("AutoPremiumHost")]
        AutoPremiumHost,

        [Description("PBIE_Azure")]
        PBIEAzure,

        [Description("Premium")]
        Premium,
    }
}
