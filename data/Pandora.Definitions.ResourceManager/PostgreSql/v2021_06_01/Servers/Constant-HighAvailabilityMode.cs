using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.PostgreSql.v2021_06_01.Servers;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum HighAvailabilityModeConstant
{
    [Description("Disabled")]
    Disabled,

    [Description("ZoneRedundant")]
    ZoneRedundant,
}